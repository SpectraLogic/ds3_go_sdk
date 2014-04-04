package net

import (
    "fmt"
    "time"
    "net/http"
    "crypto/hmac"
    "crypto/sha1"
    "encoding/base64"
    "ds3/models"
)

func setRequestHeaders(httpRequest *http.Request, creds Credentials, request models.Ds3Request) {
    now := time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05 +0000")
    httpRequest.Header.Add("Date", now)
    httpRequest.Header.Add("Authorization", buildAuthHeaderValue(creds, signatureFields{
        Verb: request.Verb().String(),
        Path: request.Path(),
        Date: now,
    }))
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
    stringToSign := fmt.Sprintf(
        "%s\n%s\n%s\n%s\n%s%s",
        fields.Verb,
        fields.ContentMd5,
        fields.ContentType,
        fields.Date,
        fields.CanonicalizedAmzHeaders,
        fields.Path,
    )
    signature := base64.StdEncoding.EncodeToString(getSha1Mac(creds.Key, stringToSign))
    return fmt.Sprintf("AWS %s:%s", creds.AccessId, signature)
}

func getSha1Mac(key, stringToSign string) []byte {
    mac := hmac.New(sha1.New, []byte(key))
    mac.Write([]byte(stringToSign))
    return mac.Sum(nil)
}

