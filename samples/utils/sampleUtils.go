// Copyright 2014-2017 Spectra Logic Corporation. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License"). You may not use
// this file except in compliance with the License. A copy of the License is located at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// or in the "license" file accompanying this file.
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package utils

import (
    "io/ioutil"
    "io"
    "bytes"
    "errors"
    "fmt"
    "github.com/SpectraLogic/ds3_go_sdk/ds3/models"
    "github.com/SpectraLogic/ds3_go_sdk/ds3"
    "math/rand"
    "time"
)

const BucketName = "GoPutBulkBucket"
const ResourceFolder = "./samples/resources/"
var BookNames = []string{"beowulf.txt", "sherlock_holmes.txt", "tale_of_two_cities.txt", "ulysses.txt"}

const PerformanceFilePrefix = "perffile_"

type PerformanceInterval struct {
    Seconds	        float64	`json:"seconds"`
    Bytes			int64	`json:"bytes"`
    BytesPerSecond	int64	`json:"bytes_per_second"`
}

type PerformanceTest struct {
    Name		string			        `json:"name"`
    NumFiles    int                     `json:"numFiles"`
    BlockSize   int64                   `json:"blockSize"`
    Start 		int64		            `json:"start"`
    Intervals 	[]PerformanceInterval	`json:"intervals"`
    Error 		string			        `json:"error"`
}

type PerformanceOutput struct {
    Put     *PerformanceTest `json:"put"`
    Get     *PerformanceTest `json:"get"`
    Errors   []string        `json:"errors"`
}

func (o *PerformanceOutput) AddError(message string) *PerformanceOutput {
    if o.Errors == nil {
        o.Errors = []string{}
    }
    o.Errors = append(o.Errors, message)
    return o
}

func NewPerformanceTest(numFiles int, blockSize int64) (*PerformanceTest) {
    return &PerformanceTest{
        Name: "",
        NumFiles: numFiles,
        BlockSize: blockSize,
        Start: time.Now().Unix(),
        Intervals: []PerformanceInterval{},
    }
}

func (p *PerformanceTest) AddInterval(sec float64, size int64) *PerformanceTest {
    int := PerformanceInterval{sec, size, int64(float64(size)/sec)}
    p.Intervals = append(p.Intervals, int)
    return p
}

// ReaderWithSizeDecorator that vends random bytes.
func BuildPerformanceReaderWithSizeDecorator(length int64) models.ReadCloserWithSizeDecorator {
    return &performanceReaderWithSizeDecorator{
        length,
    }
}

type performanceReaderWithSizeDecorator struct {
    size int64
}

func (byteReaderWithSizeDecorator *performanceReaderWithSizeDecorator) Read(b []byte) (int, error) {
    return rand.Read(b)
}

func (performanceReaderWithSizeDecorator) Close() error {
    return nil
}

func (byteReaderWithSizeDecorator *performanceReaderWithSizeDecorator) Seek(offset int64, whence int) (int64, error) {
    return offset, nil
}

func (byteReaderWithSizeDecorator *performanceReaderWithSizeDecorator) Size() (int64, error) {
    return byteReaderWithSizeDecorator.size, nil
}

// Loads a book from resources folder.
func LoadBook(book string) (models.ReaderWithSizeDecorator, error) {
    content, err := ioutil.ReadFile(ResourceFolder + book)
    if err != nil {
        return nil, err
    }

    reader := ds3.BuildByteReaderWithSizeDecorator(content)
    return reader, nil
}

// Converts a string pointer into a string. If null, returns empty. Used to
// print string pointers.
func ToSafeString(str *string) string {
    if str == nil {
        return ""
    }
    return *str
}

// Verifies the content of a ReadCloser matches the content of the specified book.
func VerifyBookContent(bookName string, content io.ReadCloser) error {
    expected, err := ioutil.ReadFile(ResourceFolder + bookName)
    if err != nil {
        return err
    }

    actual, err := ioutil.ReadAll(content)
    if err != nil {
        return err
    }

    if bytes.Compare(expected, actual) != 0 {
        errors.New(fmt.Sprintf("Mismatched content for expected and received object. Expected size %d, actual size %d.",
            len(expected),
            len(actual)))
    }
    return nil
}
