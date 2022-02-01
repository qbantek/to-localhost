package port

import (
	"fmt"
	"strconv"
)

const (
	MIN_PORT = 1
	MAX_PORT = 65535
)

type Port struct {
	Port int
}

type portInvalidError struct {
	port string
}

type portRangeError struct {
	port int
}

func (e *portInvalidError) Error() string {
	return fmt.Sprintf("invalid port: %s", e.port)
}

func (e *portRangeError) Error() string {
	return fmt.Sprintf("port %d is out of range [%d, %d]", e.port, MIN_PORT, MAX_PORT)
}

func NewPort(s string) (*Port, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return nil, fmt.Errorf("NewPort: %v", &portInvalidError{port: s})
	}
	if i < MIN_PORT || i > MAX_PORT {
		return nil, fmt.Errorf("NewPort: %v", &portRangeError{port: i})
	}

	return &Port{Port: i}, nil
}
