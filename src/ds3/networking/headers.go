package networking

import (
    "fmt"
    "time"
    "net/http"
    "crypto/hmac"
    "crypto/sha1"
    "encoding/base64"
)

func setRequestHeaders(httpRequest *http.Request, ds3Request Ds3Request, authHeaderVal string, dateHeaderVal string) (error) {
    httpRequest.Header.Add("Date", dateHeaderVal)
    httpRequest.Header.Add("Authorization", authHeaderVal)

    // Copy the headers from the request object.
    for key, val := range *ds3Request.Header() {
        httpRequest.Header.Add(key, val[0])
    }
    return nil
}

// Retrieves the current UTC date in a specific format to be used in creating Date header value
func getCurrentTime() (string) {
    return time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05 +0000")
}

type signatureFields struct {
    Verb                    string
    ContentHash             string
    ContentType             string
    Date                    string
    CanonicalizedAmzHeaders string
    Path                    string
}

func buildAuthHeaderValue(creds Credentials, fields signatureFields) string {
    // Build the string that we need to compute the MAC on.
    stringToSign := fmt.Sprintf(
        "%s\n%s\n%s\n%s\n%s%s",
        fields.Verb,
        fields.ContentHash,
        fields.ContentType,
        fields.Date,
        fields.CanonicalizedAmzHeaders,
        fields.Path,
    )

    // Compute the MAC and base64 encode the result.
    signature := computeChecksum(creds.Key, stringToSign)

    // Return the pieces in the correct format.
    return fmt.Sprintf("AWS %s:%s", creds.AccessId, signature)
}

func computeChecksum(key string, stringToSign string) (string) {
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

