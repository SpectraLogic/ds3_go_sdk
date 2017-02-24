package networking

import (
    "fmt"
    "time"
    "net/http"
    "crypto/hmac"
    "crypto/sha1"
    "encoding/base64"
)

func setRequestHeaders(httpRequest *http.Request, creds Credentials, ds3Request Ds3Request) (error) {
    // We need the current UTC date in a very specific format to be compliant.
    now := time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05 +0000")
    httpRequest.Header.Add("Date", now)

    verb, verbErr := ds3Request.Verb().String()
    if verbErr != nil {
        return verbErr
    }
    // For now we're not setting all of the signature field values because we don't need them all.
    httpRequest.Header.Add("Authorization", buildAuthHeaderValue(creds, signatureFields{
        Verb: verb,
        Path: ds3Request.Path(),
        Date: now,
    }))

    // Copy the headers from the request object.
    for key, val := range *ds3Request.Header() {
        httpRequest.Header.Add(key, val[0])
    }
    return nil
}

type signatureFields struct {
    Verb string
    ContentMd5 string
    ContentType string
    Date string
    CanonicalizedAmzHeaders string
    Path string
}

func buildAuthHeaderValue(creds Credentials, fields signatureFields) string {
    // Build the string that we need to compute the MAC on.
    stringToSign := fmt.Sprintf(
        "%s\n%s\n%s\n%s\n%s%s",
        fields.Verb,
        fields.ContentMd5,
        fields.ContentType,
        fields.Date,
        fields.CanonicalizedAmzHeaders,
        fields.Path,
    )

    // Compute the MAC and base64 encode the result.
    signature := base64.StdEncoding.EncodeToString(getSha1Mac(creds.Key, stringToSign))

    // Return the pieces in the correct format.
    return fmt.Sprintf("AWS %s:%s", creds.AccessId, signature)
}

func getSha1Mac(key, stringToSign string) []byte {
    // Create a new mac with a secret key.
    mac := hmac.New(sha1.New, []byte(key))

    // Write the string to it.
    mac.Write([]byte(stringToSign))

    // Compute the checksum.
    return mac.Sum(nil)
}

