package message

import (
	header "github.com/Pois-Noir/Botzilla-Utils/header"
)

type Message struct { // Message holds header and payload
	Header  header.Header
	Payload map[string]interface{}
}

// func NewMessage(statusCode uint8, operationCode uint8, message map[string]interface{}) *Message {
// 	return &Message{
// 		Header:  *header.NewHeader(statusCode, operationCode),
// 		Payload: message,
// 	}
// }

// func NewUserMessage(message map[string]string) *Message {

// 	message := &Message{}

// 	message.header = header.NewHeader(global_configs.OK_STATUS, global_configs.USER_MESSAGE_OPERATION_CODE)

// }

// func NewEmptyMessage() *Message {
// 	return &Message{}
// }

// func (m *Message) SetStatusCode(statusCode uint8) {
// 	m.Header.Status = statusCode
// }

// func (m *Message) SetOperationCode(operationCode uint8) {
// 	m.Header.OperationCode = operationCode
// }

// TODO
// Amir Look at this
// func (m *Message) Encode() ([]byte, error) {
// 	var messageBytes bytes.Buffer

// 	// headerBytes, err := m.Header.Encode()
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	// messageBytes.Write(headerBytes)

// 	payloadBytes, err := encoder_package.NewEncoder().EncodeMap(m.Payload)
// 	if err != nil {
// 		return nil, err
// 	}

// 	m.Header.Length = uint32(len(payloadBytes))
// 	headerBytes, err := m.Header.Encode()
// 	if err != nil {
// 		return nil, err
// 	}
// 	messageBytes.Write(headerBytes)
// 	messageBytes.Write(payloadBytes)

// 	return messageBytes.Bytes(), nil
// }

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
