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

package performance

import (
    "bytes"
    "github.com/SpectraLogic/ds3_go_sdk/ds3/models"
    "math/rand"
    "time"
)

const PerformanceFilePrefix = "perffile_"

var Randbuf = make([]byte, 1024 * 1024)

type PerformanceInterval struct {
    Seconds	        float64	`json:"seconds"`
    Bytes			int64	`json:"bytes"`
    BytesPerSecond	int64	`json:"bytes_per_second"`
}

type PerformanceBurst struct {
    ThreadID        int                     `json:"id"`
    Intervals       []PerformanceInterval   `json:"intervals"`
    Error           string                  `json:"error,omitempty"`
}

type PerformanceTest struct {
    Name		string			        `json:"name"`
    NumThreads  int                     `json:"numThreads"`
    NumFiles    int                     `json:"filesPerThread"`
    FileSize    int64                   `json:"fileSize"`
    Start 		int64		            `json:"start"`
    Seconds 	float64		            `json:"seconds"`
    Bursts 	    []PerformanceBurst	    `json:"bursts"`
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

func NewPerformanceTest(numThreads int, numFiles int, fileSize int64) (*PerformanceTest) {
    return &PerformanceTest{
        Name: "",
        NumThreads: numThreads,
        NumFiles: numFiles,
        FileSize: fileSize,
        Start: time.Now().Unix(),
        Seconds: 0.0,
        Bursts: []PerformanceBurst{},
    }
}

func NewPerformanceInterval(seconds float64, bytes int64) *PerformanceInterval {
    return &PerformanceInterval{
        Seconds: seconds,
        Bytes: bytes,
        BytesPerSecond: int64(float64(bytes)/seconds),
    }
}

func NewPerformanceBurst(threadId int) *PerformanceBurst {
    return &PerformanceBurst{
        ThreadID: threadId,
        Intervals: []PerformanceInterval{},
    }
}

func NewPerformanceBurstError(err string) *PerformanceBurst {
    return &PerformanceBurst{Error: err,}
}


func (p *PerformanceBurst) AddInterval(sec float64, size int64) *PerformanceBurst {
    interval := NewPerformanceInterval(sec, size)
    p.Intervals = append(p.Intervals, *interval)
    return p
}

// ReaderWithSizeDecorator that vends random bytes.
func BuildPerformanceReaderWithSizeDecorator(length int64) models.ReadCloserWithSizeDecorator {
    rand.Read(Randbuf)
    return &performanceReaderWithSizeDecorator{
        length,
    }
}

type performanceReaderWithSizeDecorator struct {
    size int64
}

func (byteReaderWithSizeDecorator *performanceReaderWithSizeDecorator) Read(b []byte) (int, error) {
    return bytes.NewBuffer(Randbuf).Read(b)
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
