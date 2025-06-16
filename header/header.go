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
	PayloadLength uint32 //4 bytes payload length int
}

func NewHeader(status uint8, operationCode uint8, len uint32) *Header {
	return &Header{
		Status:        status,
		OperationCode: operationCode,
		PayloadLength: len,
	}
}

func (h *Header) Encode() []byte {

	buf := [global_configs.HEADER_LENGTH]byte{}

	buf[global_configs.STATUS_CODE_INDEX] = h.Status
	buf[global_configs.OPERATION_CODE_INDEX] = h.OperationCode

	binary.BigEndian.PutUint32(buf[global_configs.MESSAGE_LENGTH_INDEX:], h.PayloadLength)

	return buf[:]

}

func Decode(buffer []byte) (*Header, error) {

	if len(buffer) != global_configs.HEADER_LENGTH {
		return nil, errors.New("buffer size is not correct")
	}

	header := &Header{}

	header.OperationCode = uint8(buffer[global_configs.STATUS_CODE_INDEX])
	header.Status = uint8(buffer[global_configs.OPERATION_CODE_INDEX])
	header.PayloadLength = binary.BigEndian.Uint32(buffer[global_configs.MESSAGE_LENGTH_INDEX:global_configs.HEADER_LENGTH])

	return header, nil

}
