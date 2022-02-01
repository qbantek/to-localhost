package port_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/qbantek/to-localhost/internal/port"
)

func TestNewPort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	n := randomInteger(1, 65535)

	cases := []struct {
		name         string
		port         string
		want         int
		errorMessage string
	}{
		{fmt.Sprintf("valid port:%d", n), fmt.Sprintf("%d", n), n, ""},
		{"invalid port:too_small", "0", 0, "NewPort: port 0 is out of range [1, 65535]"},
		{"invalid port:too_big", "65536", 0, "NewPort: port 65536 is out of range [1, 65535]"},
		{"invalid port:non_numeric", "foo", 0, "NewPort: invalid port: foo"},
		{"invalid port:empty_string", "", 0, "NewPort: invalid port: "},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := port.NewPort(c.port)
			if err != nil {
				if err.Error() != c.errorMessage {
					t.Errorf("Expected error: %s, got: %s", c.errorMessage, err.Error())
				}
				return
			}

			if c.errorMessage != "" {
				t.Errorf("Expected error: %s", c.errorMessage)
				return
			}

			if got.Port != c.want {
				t.Errorf("got: %d, want: %d", got.Port, c.want)
			}
		})
	}
}

func randomInteger(min, max int) int {
	return rand.Intn(max-min) + min
}
