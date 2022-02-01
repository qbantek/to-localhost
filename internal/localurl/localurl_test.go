package localurl_test

import (
	"testing"

	"github.com/qbantek/to-localhost/internal/localurl"
)

func TestNewURL(t *testing.T) {
	cases := []struct {
		name         string
		port         string
		want         string
		errorMessage string
	}{
		{"Default HTTP Port", "80", "http://localhost", ""},
		{"A valid Port", "8080", "http://localhost:8080", ""},
		{"Invalid Port value", "foo", "", "NewURL: NewPort: invalid port: foo"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := localurl.NewURL(c.port)
			if err != nil {
				if err.Error() != c.errorMessage {
					t.Errorf("Expected error: %s, got: %s", c.errorMessage, err.Error())
				}
				return
			}

			if c.errorMessage != "" {
				t.Errorf("Expected error: %s", c.errorMessage)
			}

			if got.String() != c.want {
				t.Errorf("NewURL(%s) got %s, want %s", c.port, got, c.want)
			}
		})
	}
}
