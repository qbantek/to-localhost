package localhosturl

import (
	"fmt"
	"strconv"
)

const (
	MIN_PORT      = 1
	MAX_PORT      = 65535
	LOCALHOST_URL = "http://localhost"
)

type portError struct {
	port string
}

func (e *portError) Error() string {
	return fmt.Sprintf("Invalid port value: [%s]", e.port)
}

func NewURL(port string) (string, error) {
	// HTTP default port is 80
	if port == "80" {
		return LOCALHOST_URL, nil
	}

	err := validatePort(port)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s:%s", LOCALHOST_URL, port), nil
}

func validatePort(port string) error {
	portInt, err := strconv.Atoi(port)
	if err != nil || portInt < MIN_PORT || portInt > MAX_PORT {
		return &portError{port}
	}

	return nil
}
