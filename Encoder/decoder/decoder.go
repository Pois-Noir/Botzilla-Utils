package decoder

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math"
)

const (
	TypeString  = 0x01
	TypeInt64   = 0x02
	TypeFloat64 = 0x03
	TypeBool    = 0x04
	TypeMap     = 0x05
	TypeSlice   = 0x06
)

type Decoder struct {
	Data      []byte
	bytesRead int
}

func NewDecoder(data []byte) *Decoder {
	return &Decoder{
		Data: data,
	}
}

func (d *Decoder) Decode() (map[string]interface{}, error) {

	msgLen := len(d.Data)
	result := make(map[string]interface{})
	startBytes := d.bytesRead
	for d.bytesRead-startBytes < msgLen {
		key, err := d.readString()
		if err != nil {
			return nil, err
		}
		val, err := d.readValue()
		if err != nil {
			return nil, err
		}
		result[key] = val
	}
	return result, nil
}

func (d *Decoder) readType() (byte, error) {
	var t [1]byte
	if err := d.readFull(t[:]); err != nil {
		return 0, err
	}
	return t[0], nil
}

func (d *Decoder) readUint16() (uint16, error) {
	var buf [2]byte
	if err := d.readFull(buf[:]); err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint16(buf[:]), nil
}

func (d *Decoder) readBytes(n uint16) ([]byte, error) {
	b := make([]byte, n)
	if err := d.readFull(b); err != nil {
		return nil, err
	}
	return b, nil
}

func (d *Decoder) readString() (string, error) {
	t, err := d.readType()
	if err != nil {
		return "", err
	}
	if t != TypeString {
		return "", fmt.Errorf("expected string type, got 0x%02x", t)
	}
	length, err := d.readUint16()
	if err != nil {
		return "", err
	}
	data, err := d.readBytes(length)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (d *Decoder) readValue() (interface{}, error) {
	t, err := d.readType()
	if err != nil {
		return nil, err
	}
	length, err := d.readUint16()
	if err != nil {
		return nil, err
	}
	payload, err := d.readBytes(length)
	if err != nil {
		return nil, err
	}

	switch t {
	case TypeString:
		return string(payload), nil
	case TypeInt64:
		if len(payload) != 8 {
			return nil, errors.New("invalid int64 length")
		}
		return int64(binary.BigEndian.Uint64(payload)), nil
	case TypeFloat64:
		if len(payload) != 8 {
			return nil, errors.New("invalid float64 length")
		}
		bits := binary.BigEndian.Uint64(payload)
		return math.Float64frombits(bits), nil
	case TypeBool:
		if len(payload) != 1 {
			return nil, errors.New("invalid bool length")
		}
		return payload[0] == 1, nil
	case TypeMap:
		return decodeNestedMap(payload)
	case TypeSlice:
		return decodeNestedSlice(payload)
	default:
		return nil, fmt.Errorf("unsupported type marker 0x%02x", t)
	}
}

// readFull wraps io.ReadFull and tracks bytes read
func (d *Decoder) readFull(buf []byte) error {
	n, err := io.ReadFull(d.reader, buf)
	d.bytesRead += n
	return err
}

// Nested decoding helpers
func decodeNestedMap(data []byte) (map[string]interface{}, error) {
	subDecoder := NewDecoderBytes(data)
	return subDecoder.Decode(len(data))
}

func decodeNestedSlice(data []byte) ([]interface{}, error) {
	subDecoder := NewDecoderBytes(data)
	var result []interface{}
	for {
		val, err := subDecoder.readValue()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		result = append(result, val)
	}
	return result, nil
}
