package payload

type Payload interface {
	Encode() []byte
	Decode([]byte)
}

type Json map[string]string

type RegisterComponent struct {
	Name string
	Port uint32
}
