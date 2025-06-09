package message

import (
    "encoding/binary"
    "encoding/json"
    "fmt"
    "io"
)

type Message struct { // Message holds header and payload
    Header  Header
    Payload map[string]interface{}
}

/*
func (m *Message) Encode() ([]byte, error) {
    // 1) marshal payload
    body, err := json.Marshal(m.Payload)
    if err != nil {
        return nil, fmt.Errorf("Message.Encode: json.Marshal: %w", err)
    }
	
	encoder := NewEncoder() 
	byteStrean := encoder.EncoderMsg(m)

    // 2) sync header length
    m.Header.Length = uint32(len(body))

    // 3) encode header
    hdr := m.Header.Encode() // [5]byte

    // 4) concat and return
    packet := append(hdr[:], body...)
    return packet, nil
}*/

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
