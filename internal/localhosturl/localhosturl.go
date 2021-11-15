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
	return fmt.Sprintf("Invalid port value: %s", e.port)
}

func NewURL(port string) (string, error) {
	url := LOCALHOST_URL

	if port != "" && port != "80" {
		portInt, err := strconv.Atoi(port)
		if err != nil || portInt < MIN_PORT || portInt > MAX_PORT {
			return "", &portError{port}
		}
		url += ":" + port
	}

	return url, nil
}
