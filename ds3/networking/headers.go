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

package networking

import (
    "fmt"
    "time"
    "crypto/hmac"
    "crypto/sha1"
    "encoding/base64"
    "errors"
    "github.com/SpectraLogic/ds3_go_sdk/ds3/models"
    "net/url"
)

// Http Headers
const ContentCRC32 string = "Content-CRC32"
const ContentCRC32C string = "Content-CRC32C"
const ContentMd5 string = "Content-MD5"
const ContentSha256 string = "Content-SHA256"
const ContentSha512 string = "Content-SHA512"
const ContentTypeKey string = "Content-Type" // This header describes the format of the contents

type headerFields struct {
    AuthHeaderVal string
    DateHeaderVal string
    ChecksumType  models.ChecksumType
    ContentHash   string
}

// Gets the correct header key for the specified checksum type
func getChecksumHeaderKey(checksumType models.ChecksumType) (string, error) {
    switch checksumType {
        case models.CHECKSUM_TYPE_CRC_32: return ContentCRC32, nil
        case models.CHECKSUM_TYPE_CRC_32C: return ContentCRC32C, nil
        case models.CHECKSUM_TYPE_MD5: return ContentMd5, nil
        case models.CHECKSUM_TYPE_SHA_256: return ContentSha256, nil
        case models.CHECKSUM_TYPE_SHA_512: return ContentSha512, nil
        case models.NONE: return "", nil
        default:
            return "", errors.New(fmt.Sprintf("Invalid ChecksumType represented by: %d", checksumType))
    }
}

// Retrieves the current UTC date in a specific format to be used in creating the Date header value
func getCurrentTime() (string) {
    return time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05 +0000")
}

type signatureFields struct {
    Verb                      string
    ContentHash               string
    ContentType               string
    Date                      string
    CanonicalizedAmzHeaders   string
    CanonicalizedSubResources string
    Path                      string
}

func (fields *signatureFields) BuildAuthHeaderValue(creds *Credentials) (string) {
    // Build the string that we need to compute the MAC on.

    // Create a URL containing just the path so that we can use the library's escaping logic
    // to match the http request escaping
    var urlPath url.URL
    urlPath.Path = fields.Path

    stringToSign := fmt.Sprintf(
        "%s\n%s\n%s\n%s\n%s%s%s",
        fields.Verb,
        fields.ContentHash,
        fields.ContentType,
        fields.Date,
        fields.CanonicalizedAmzHeaders,
        urlPath.EscapedPath(),
        fields.CanonicalizedSubResources,
    )

    signature := computeSignature(creds.Key, stringToSign)

    // Return the pieces in the correct format.
    return fmt.Sprintf("AWS %s:%s", creds.AccessId, signature)
}

func computeSignature(key string, stringToSign string) (string) {
    // Compute the MAC and base64 encode the result.
    return base64.StdEncoding.EncodeToString(calcSha1Mac(key, stringToSign))
}

func calcSha1Mac(key, stringToSign string) []byte {
    // Create a new mac with a secret key.
    mac := hmac.New(sha1.New, []byte(key))

    // Write the string to it.
    mac.Write([]byte(stringToSign))

    // Compute the checksum.
    return mac.Sum(nil)
}

