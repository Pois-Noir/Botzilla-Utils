package header

import (
    "encoding/binary"
    "encoding/json"
    "fmt"
    "io"
)

// pipe & filter style (dont care where its coming from and where it is going)
//Header is the 5-byte prefix
type Header struct {
    Status uint8 //1 byte status code uint8
    Length uint32 //4 bytes payload length int
}