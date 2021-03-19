package helpers

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"github.com/SpectraLogic/ds3_go_sdk/ds3/buildclient"
	ds3Models "github.com/SpectraLogic/ds3_go_sdk/ds3/models"
	"github.com/SpectraLogic/ds3_go_sdk/ds3_utils/ds3Testing"
	"github.com/SpectraLogic/ds3_go_sdk/helpers/channels"
	"github.com/SpectraLogic/ds3_go_sdk/helpers/models"
	"hash"
	"io"
	"reflect"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type randomDataReader struct {
	curRead int
	maxRead int
	mutex sync.Mutex
	totalFileHash hash.Hash
	rangesHash hash.Hash
	ranges []ds3Models.Range
}

func (reader *randomDataReader) Read(p []byte) (int, error) {
	reader.mutex.Lock()
	defer reader.mutex.Unlock()

	if reader.totalFileHash == nil {
		reader.totalFileHash = sha256.New()
	}
	if reader.rangesHash == nil {
		reader.rangesHash = sha256.New()
	}

	if reader.curRead >= reader.maxRead {
		return 0, io.EOF
	}
	amountToRead := reader.maxRead - reader.curRead
	if amountToRead > len(p) {
		amountToRead = len(p)
	}

	buff := make([]byte, amountToRead)
	_, _ = rand.Read(buff)
	reader.totalFileHash.Write(buff)
	for i := 0; i < len(buff); i++ {
		p[i] = buff[i]

		// calculate the hash on specified ranges
		curByteOffset := int64(reader.curRead + i)

		for _, curRange := range reader.ranges {
			if curByteOffset >= curRange.Start && curByteOffset <= curRange.End {
				reader.rangesHash.Write([]byte{buff[i]})
				break
			}
		}
	}

	reader.curRead += amountToRead
	return amountToRead, nil
}

func (reader *randomDataReader) Close() error {
	return nil
}

type randomDataReadChannelBuilder struct {
	randomDataReader *randomDataReader
	channels.FatalErrorHandler
}

func (builder *randomDataReadChannelBuilder) IsChannelAvailable(offset int64) bool {
	 return int64(builder.randomDataReader.curRead) == offset
}

func (builder *randomDataReadChannelBuilder) GetChannel(_ int64) (io.ReadCloser, error) {
	return builder.randomDataReader, nil
}

func (builder *randomDataReadChannelBuilder) OnDone(reader io.ReadCloser) {
	_ = reader.Close()
}

type consumeWriter struct {
	size int64
	hash hash.Hash
	mutex sync.Mutex
}

func (writer *consumeWriter) Write(p []byte) (n int, err error) {
	writer.mutex.Lock()
	defer writer.mutex.Unlock()

	if writer.hash == nil {
		writer.hash = sha256.New()
	}
	atomic.AddInt64(&writer.size, int64(len(p)))
	writer.hash.Write(p)
	return len(p), err
}

func (writer *consumeWriter) Close() error {
	return nil
}

type consumerWriteChannelBuilder struct {
	writer *consumeWriter
	channels.FatalErrorHandler
}

func (builder *consumerWriteChannelBuilder) IsChannelAvailable(_ int64) bool {
	return true
}

func (builder *consumerWriteChannelBuilder) GetChannel(_ int64) (io.WriteCloser, error) {
	return builder.writer, nil
}

func (builder *consumerWriteChannelBuilder) OnDone(writer io.WriteCloser) {
	_ = writer.Close()
}

func TestRetrievingObjectLargerThanCacheInOrder(t *testing.T) {
	t.Skip("Test reduces BP cache size to 1 GB and does not reset it. Best run against the BP simulator.")

	const minCache = 1073741824
	const objectSize = minCache + 100

	testRanges := []ds3Models.Range{{Start: 2, End: 10}, {Start: 30, End: 40}, {Start: 50, End: 60}, {Start: 80, End: objectSize-1}}

	bucketName := fmt.Sprintf("testBucket-%d", time.Now().UnixNano())
	objectName := fmt.Sprintf("testObject-%d", time.Now().UnixNano())

	simClient, err := buildclient.FromEnv()
	ds3Testing.AssertNilError(t, err)

	// Ensure the cache is set to minimum of 1 GB
	getCacheResp, err := simClient.GetCacheFilesystemsSpectraS3(ds3Models.NewGetCacheFilesystemsSpectraS3Request())
	ds3Testing.AssertNilError(t, err)

	t.Logf("CACHE:")
	for i, curCache := range getCacheResp.CacheFilesystemList.CacheFilesystems {
		t.Logf("%d) %s, (%v)", i, curCache.Id, curCache.MaxCapacityInBytes)

		req := ds3Models.NewModifyCacheFilesystemSpectraS3Request(curCache.Id).WithMaxCapacityInBytes(minCache)
		modifyResp, err := simClient.ModifyCacheFilesystemSpectraS3(req)
		ds3Testing.AssertNilError(t, err)
		t.Logf("modified: %s, (%v)", modifyResp.CacheFilesystem.Id, modifyResp.CacheFilesystem.MaxCapacityInBytes)
	}

	// Make sure that our test bucket exists with a file in it
	_, err = simClient.GetBucketSpectraS3(ds3Models.NewGetBucketSpectraS3Request(bucketName))
	if err != nil {
		t.Logf("Creating bucket: %s", bucketName)
		_, err := simClient.PutBucket(ds3Models.NewPutBucketRequest(bucketName))
		ds3Testing.AssertNilError(t, err)

		defer func() {
			_, err := simClient.DeleteBucketSpectraS3(ds3Models.NewDeleteBucketSpectraS3Request(bucketName).WithForce())
			if err != nil {
				t.Errorf("failed to delete bucket with force '%s': %v", bucketName, err)
			}
		}()
	}

	// Create the object
	t.Logf("Creating object: %s", objectName)

	helperWrapper := NewHelpers(simClient)

	objReader := randomDataReader{
		maxRead: objectSize,
		ranges: testRanges,
	}

	putObject := models.PutObject{
		PutObject:      ds3Models.Ds3PutObject{
			Name: objectName,
			Size: int64(objectSize),
		},
		ChannelBuilder: &randomDataReadChannelBuilder{
			randomDataReader:  &objReader,
			FatalErrorHandler: channels.FatalErrorHandler{},
		},
	}

	writeStrategy := WriteTransferStrategy{
		BlobStrategy: &SimpleBlobStrategy{
			Delay: time.Millisecond * 10,
			MaxConcurrentTransfers: 5,
			MaxWaitingTransfers: 1,
		},
		Options:      WriteBulkJobOptions{MaxUploadSize: &MinUploadSize},
		Listeners:    ListenerStrategy{
			ErrorCallback: func(objectName string, err error) {
				t.Errorf("unexpected error on '%s': %v", objectName, err)
			},
		},
	}

	_, err = helperWrapper.PutObjects(bucketName, []models.PutObject{putObject}, writeStrategy)
	ds3Testing.AssertNilError(t, err)

	defer func() {
		_, err := simClient.DeleteObject(ds3Models.NewDeleteObjectRequest(bucketName, objectName))
		if err != nil {
			t.Errorf("failed to delete object %s: %v", objectName, err)
		}
	}()

	testCases := []struct{
		name string
		ranges []ds3Models.Range
	} {
		{ name: "retrieve entire file", ranges: nil },
		{ name: "retrieve ranges", ranges: testRanges},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// reclaim cache since simulator doesn't seem to do it on its own
			_, err := simClient.ForceFullCacheReclaimSpectraS3(ds3Models.NewForceFullCacheReclaimSpectraS3Request())
			ds3Testing.AssertNilError(t, err)

			readStrategy := ReadTransferStrategy{
				BlobStrategy: &SimpleBlobStrategy{
					Delay: time.Millisecond * 10,
					MaxConcurrentTransfers: 5,
					MaxWaitingTransfers: 1,
				},
				Options:      ReadBulkJobOptions{ChunkClientProcessingOrderGuarantee: ds3Models.JOB_CHUNK_CLIENT_PROCESSING_ORDER_GUARANTEE_IN_ORDER},
				Listeners:    ListenerStrategy{
					ErrorCallback: func(objectName string, err error) {
						t.Errorf("unexpected error on '%s': %v", objectName, err)
					},
				},
			}

			objWriter := consumeWriter{}
			getObject := models.GetObject{
				Name:   objectName,
				Ranges: testCase.ranges,
				ChannelBuilder: &consumerWriteChannelBuilder{
					writer: &objWriter,
					FatalErrorHandler: channels.FatalErrorHandler{},
				},
			}
			getJobIds, err := helperWrapper.GetObjectsSpanningJobs(bucketName, []models.GetObject{getObject}, readStrategy)
			ds3Testing.AssertNilError(t, err)

			// verify amount of data retrieved
			if len(testCase.ranges) == 0 {
				ds3Testing.AssertInt64(t, "bytes retrieved", objectSize, objWriter.size)
			} else {
				expectedDataRetrieved := int64(0)
				for _, curRange := range testCase.ranges {
					expectedDataRetrieved += curRange.End - curRange.Start + 1
				}
				ds3Testing.AssertInt64(t, "bytes retrieved", expectedDataRetrieved, objWriter.size)
			}

			// verify retrieved data hash
			var readHash []byte
			writeHash := objWriter.hash.Sum(nil)
			if len(testCase.ranges) == 0 {
				readHash = objReader.totalFileHash.Sum(nil)
			} else {
				readHash = objReader.rangesHash.Sum(nil)
			}
			ds3Testing.AssertBool(t, "hash matches", true, reflect.DeepEqual(readHash, writeHash))
			ds3Testing.AssertBool(t, "multiple job ids returned", true, len(getJobIds) > 1)
			t.Logf("Get Job: %v", getJobIds)

			ds3Testing.AssertBool(t, "has fatal err", false, getObject.ChannelBuilder.HasFatalError())
		})
	}
}
