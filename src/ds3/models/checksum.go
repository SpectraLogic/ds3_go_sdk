package models

type Checksum struct {
    ContentHash string
    Type        ChecksumType
}

func NewNoneChecksum() Checksum {
    return Checksum{
        Type: NONE,
        ContentHash: "" }
}
