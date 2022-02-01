package localurl

import (
	"fmt"

	"github.com/qbantek/to-localhost/internal/port"
)

const (
	LOCALHOST_URL = "http://localhost"
)

type LocalURL struct {
	port int
}

func (u *LocalURL) String() string {
	if u.port == 80 {
		return LOCALHOST_URL
	}
	return fmt.Sprintf("%s:%d", LOCALHOST_URL, u.port)
}

func NewURL(sPort string) (*LocalURL, error) {
	p, err := port.NewPort(sPort)
	if err != nil {
		return nil, fmt.Errorf("NewURL: %s", err)
	}

	return &LocalURL{port: p.Port}, nil
}
