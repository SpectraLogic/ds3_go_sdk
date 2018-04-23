package models

// Describes a blob
type BlobDescription struct {
    name   string
    offset int64
    length int64
}

func NewBlobDescription(name string, offset int64, length int64) BlobDescription {
    return BlobDescription{
        name:   name,
        offset: offset,
        length: length,
    }
}

func (d *BlobDescription) Name() string {
    return d.name
}

func (d *BlobDescription) Offset() int64 {
    return d.offset
}

func (d *BlobDescription) Length() int64 {
    return d.length
}
