package localhosturl_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/qbantek/to-localhost/internal/localhosturl"
)

func TestNewURL(t *testing.T) {
	// get a random Port number between 1 and 65535
	rand.Seed(time.Now().UnixNano())
	n := fmt.Sprintf("%d", randomInteger(1, 65535))

	cases := []struct {
		name         string
		port         string
		want         string
		errorMessage string
	}{
		{"Default HTTP Port", "80", "http://localhost", ""},
		{"A valid Port", n, "http://localhost:" + n, ""},
		{"Non numeric Port value", "invalid", "", "Invalid port value: [invalid]"},
		{"Empty String", "", "", "Invalid port value: []"},
		{"Port number < 1", "0", "", "Invalid port value: [0]"},
		{"Port number > 65535", "65536", "", "Invalid port value: [65536]"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := localhosturl.NewURL(c.port)
			if err != nil {
				if err.Error() != c.errorMessage {
					t.Errorf("NewURL(%s) = %s", c.port, err.Error())
				}
			}
			if got != c.want {
				t.Errorf("NewURL(%s) got %s, want %s", c.port, got, c.want)
			}
		})
	}
}

func randomInteger(min, max int) int {
	return rand.Intn(max-min) + min
}
