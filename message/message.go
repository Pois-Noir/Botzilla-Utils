package message

import (
	"bytes"

	header_package "github.com/Pois-Noir/Botzilla-Utils/header"
	encoder_package "github.com/Pois-Noir/Mammad/encoder"
)

type Message struct { // Message holds header and payload
	Header  header_package.Header
	Payload map[string]interface{}
}

func NewMessage(statusCode uint8, message map[string]interface{}) *Message {
	return &Message{
		Header: header_package.Header{
			Status: statusCode,
		},
		Payload: message,
	}
}

func (m *Message) SetStatusCode(statusCode uint8) {
	m.Header.Status = statusCode
}

func (m *Message) Encode() ([]byte, error) {
	var messageBytes bytes.Buffer

	// headerBytes, err := m.Header.Encode()
	// if err != nil {
	// 	return nil, err
	// }
	// messageBytes.Write(headerBytes)

	payloadBytes, err := encoder_package.NewEncoder().EncodeMap(m.Payload)
	if err != nil {
		return nil, err
	}

	m.Header.Length = uint32(len(payloadBytes))
	headerBytes, err := m.Header.Encode()
	if err != nil {
		return nil, err
	}
	messageBytes.Write(headerBytes)
	messageBytes.Write(payloadBytes)

	return messageBytes.Bytes(), nil
}

/*


/*
// Decode reads one Message from r
func Decode(r io.Reader) (*Message, error) {
    // 1) read header bytes
    var hdrBuf [5]byte
    if _, err := io.ReadFull(r, hdrBuf[:]); err != nil {
        return nil, fmt.Errorf("Decode: read header: %w", err)
    }

    // 2) parse header
    h, err := DecodeHeader(hdrBuf[:])
    if err != nil {
        return nil, fmt.Errorf("Decode: %w", err)
    }

    // 3) read payload
    body := make([]byte, h.Length)
    if _, err := io.ReadFull(r, body); err != nil {
        return nil, fmt.Errorf("Decode: read payload: %w", err)
    }

    // 4) unmarshal JSON
    var payload map[string]interface{}
    if err := json.Unmarshal(body, &payload); err != nil {
        return nil, fmt.Errorf("Decode: json.Unmarshal: %w", err)
    }

    return &Message{
        Header:  h,
        Payload: payload,
    }, nil
}*/
