package networking

type Checksum struct {
    ContentHash string
    Type        ChecksumType
}

func NewNoneChecksum() Checksum {
    return Checksum{
        Type: NONE,
        ContentHash: "" }
}

type ChecksumType int

const(
    CRC_32 ChecksumType = iota
    CRC_32C
    MD5
    SHA_256
    SHA_512
    NONE
)
