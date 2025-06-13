package header

import (
	"bytes"
	"encoding/binary"
)

// pipe & filter style (dont care where its coming from and where it is going)
// Header is the 5-byte prefix
type Header struct {
	Status        uint8  //1 byte status code uint8
	OperationCode uint8  // 1 byte OperationCode
	length        uint32 //4 bytes payload length int
}

func NewEmptyHeader() *Header {
	return &Header{}
}

func NewHeader(status uint8, operationCode uint8) *Header {
	return &Header{
		Status:        status,
		OperationCode: operationCode,
	}
}

func (h *Header) SetMessageLength(messageLen uint32) {
	h.length = messageLen
}

func (h *Header) Encode() ([]byte, error) {
	var buf bytes.Buffer

	err := buf.WriteByte(h.Status)
	if err != nil {
		return nil, err
	}

	err = buf.WriteByte(h.OperationCode)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buf, binary.BigEndian, h.length)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil

}
