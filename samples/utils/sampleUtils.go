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
)

const BucketName = "GoPutBulkBucket"
const ResourceFolder = "./samples/resources/"
var BookNames = []string{"beowulf.txt", "sherlock_holmes.txt", "tale_of_two_cities.txt", "ulysses.txt"}

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
