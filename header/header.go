package header

import (
	"bytes"
	"encoding/binary"
)

// pipe & filter style (dont care where its coming from and where it is going)
// Header is the 5-byte prefix
type Header struct {
	Status uint8  //1 byte status code uint8
	Length uint32 //4 bytes payload length int
}

func (h *Header) Encode() ([]byte, error) {
	var buf bytes.Buffer

	err := buf.WriteByte(h.Status)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buf, binary.BigEndian, h.Length)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil

}
