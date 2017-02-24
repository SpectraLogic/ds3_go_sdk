package models

type Object struct {
    ETag string
    Key string
    LastModified string
    Owner Owner
    Size int64
    StorageClass string
}

