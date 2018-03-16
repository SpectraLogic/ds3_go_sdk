// Copyright 2014-2018 Spectra Logic Corporation. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License"). You may not use
// this file except in compliance with the License. A copy of the License is located at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// or in the "license" file accompanying this file.
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package ds3_integration

import "io"

type nilReadSizer struct {
	reader io.Reader
	fileSize int64
}

func newNilReadSizer(reader io.Reader, fileSize int64) nilReadSizer {
	return nilReadSizer { reader: reader, fileSize: fileSize }
}

func (readSizer *nilReadSizer) Read(p []byte) (n int, err error) {
	return readSizer.reader.Read(p)
}

func (readSizer *nilReadSizer) Size() (int64, error) {
	return readSizer.fileSize, nil
}
