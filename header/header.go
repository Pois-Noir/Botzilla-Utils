package header

import (
	"bytes"
	"encoding/binary"
	"io"
	"net"
)

const STATUDBYTEINDEX = 0
const OPERATIONCODEBYTEINDEX = 1
const MESSAGELENGTHBYTEINDEXSTART = 2
const MESSAGEBYTECOUNT = 4

// pipe & filter style (dont care where its coming from and where it is going)
// Header is the 5-byte prefix
type Header struct {
	Status        uint8  //1 byte status code uint8
	OperationCode uint8  // 1 byte OperationCode
	Length        uint32 //4 bytes payload length int
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
	h.Length = messageLen
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

	err = binary.Write(&buf, binary.BigEndian, h.Length)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil

}

func DecodeHeader(conn net.Conn) (*Header, error) {

	var headerBytes [6]byte

	_, err := io.ReadFull(conn, headerBytes[:])
	if err != nil {
		// handle error, e.g., log or return
		return nil, err
	}

	header := NewHeader(uint8(headerBytes[STATUDBYTEINDEX]), uint8(headerBytes[OPERATIONCODEBYTEINDEX]))
	header.SetMessageLength(binary.BigEndian.Uint32(headerBytes[MESSAGELENGTHBYTEINDEXSTART : MESSAGELENGTHBYTEINDEXSTART+MESSAGEBYTECOUNT]))
	return header, nil
}
