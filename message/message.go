package message

import (
	header_package "github.com/Pois-Noir/Botzilla-Utils/header"
)

type Message struct { // Message holds header and payload
	Header  header_package.Header
	Payload map[string]interface{}
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
