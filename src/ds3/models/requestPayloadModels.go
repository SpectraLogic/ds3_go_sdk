package models

type CompleteMultipartUpload struct {
    Parts []Part `xml:"Part"`
}

type Part struct {
    PartNumber int
    ETag string
}

type Ds3PutObject struct {
    Name string `xml:"Name,attr"`
    Size int64 `xml:"Size,attr"`
}

type Ds3GetObject struct {
    Name string `xml:"Name,attr"`
    Length *int64 `xml:"Length,attr,omitempty"`
    Offset *int64 `xml:"Offset,attr,omitempty"`
}

func NewDs3GetObject(name string) Ds3GetObject {
    return Ds3GetObject{Name:name}
}

// Creates a Ds3GetObject used for partial objects
func NewPartialDs3GetObject(name string, length int64, offset int64) Ds3GetObject {
    return Ds3GetObject{
        Name:name,
        Length:&length,
        Offset:&offset,
    }
}
