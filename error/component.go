package error

import "fmt"

var (
	ErrComponentNotCreatedType *ErrComponentNotCreated
)

type ErrComponentNotCreated struct {
	serverAddr string
	name       string
	port       int
}

func (e *ErrComponentNotCreated) Error() string {
	return fmt.Sprintf("component not created: server address: %s, component name %s, component port %d", e.serverAddr, e.name, e.port)

}

func NewErrComponentNotCreated(serverAddr string, name string, port int) error {
	return &ErrComponentNotCreated{
		serverAddr: serverAddr,
		name:       name,
		port:       port,
	}
}
