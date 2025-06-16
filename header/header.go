package header

import (
	"encoding/binary"
	"errors"

	"github.com/Pois-Noir/Botzilla-Utils/global_configs"
)

// pipe & filter style (dont care where its coming from and where it is going)
// Header is the 5-byte prefix
type Header struct {
	Status        uint8  //1 byte status code uint8
	OperationCode uint8  // 1 byte OperationCode
	Length        uint32 //4 bytes payload length int
}

// func NewEmptyHeader() *Header {
// 	return &Header{}
// }

func NewHeader(status uint8, operationCode uint8, len uint32) *Header {
	return &Header{
		Status:        status,
		OperationCode: operationCode,
		Length:        len,
	}
}

// func (h *Header) SetMessageLength(messageLen uint32) {
// 	h.Length = messageLen
// }

func (h *Header) Encode() []byte {

	buf := [global_configs.HEADER_LENGTH]byte{}

	buf[global_configs.STATUS_CODE_INDEX] = h.Status
	buf[global_configs.OPERATION_CODE_INDEX] = h.OperationCode

	binary.BigEndian.PutUint32(buf[global_configs.MESSAGE_LENGTH_INDEX:], h.Length)

	return buf[:]

}

func Decode(buffer []byte) (*Header, error) {

	if len(buffer) != global_configs.HEADER_LENGTH {
		return nil, errors.New("buffer size is not correct")
	}

	header := &Header{}

	header.OperationCode = uint8(buffer[global_configs.STATUS_CODE_INDEX])
	header.Status = uint8(buffer[global_configs.OPERATION_CODE_INDEX])
	header.Length = binary.BigEndian.Uint32(buffer[global_configs.MESSAGE_LENGTH_INDEX:global_configs.HEADER_LENGTH])

	return header, nil

}

// func DecodeHeader(conn net.Conn) (*Header, error) {

// 	var headerBytes [6]byte

// 	_, err := io.ReadFull(conn, headerBytes[:])
// 	if err != nil {
// 		// handle error, e.g., log or return
// 		return nil, err
// 	}

// 	header := NewHeader(uint8(headerBytes[STATUDBYTEINDEX]), uint8(headerBytes[OPERATIONCODEBYTEINDEX]))
// header.SetMessageLength(binary.BigEndian.Uint32(headerBytes[MESSAGELENGTHBYTEINDEXSTART : MESSAGELENGTHBYTEINDEXSTART+MESSAGEBYTECOUNT]))
// 	return header, nil
// }

// func DecodeHeaderBuffered(bReader *bufio.Reader) (*Header, error) {
// 	var headerBytes [6]byte

// 	_, err := io.ReadFull(bReader, headerBytes[:])
// 	if err != nil {
// 		// handle error, e.g., log or return
// 		return nil, err
// 	}

// 	header := NewHeader(uint8(headerBytes[STATUDBYTEINDEX]), uint8(headerBytes[OPERATIONCODEBYTEINDEX]))
// 	header.SetMessageLength(binary.BigEndian.Uint32(headerBytes[MESSAGELENGTHBYTEINDEXSTART : MESSAGELENGTHBYTEINDEXSTART+MESSAGEBYTECOUNT]))
// 	return header, nil
// }
