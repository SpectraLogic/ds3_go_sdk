package models

type Enum int

func (enum Enum) String() string {
    return "something..."
}

const (
    UNDEFINED = 0
)