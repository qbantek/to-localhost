package config_test

import (
	"os"
	"testing"

	"github.com/qbantek/to-localhost/internal/config"
)

func TestNewConfig(t *testing.T) {
	cases := []struct {
		name         string
		env          string
		want         string
		errorMessage string
	}{
		{"DefaultPort", "", "5000", ""},
		{"CustomPort", "8080", "8080", ""},
		{"InvalidPort", "foo", "", "NewConfig: NewPort: invalid port: foo"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			os.Setenv("PORT", c.env)

			cfg, err := config.NewConfig()
			if err != nil {
				if err.Error() != c.errorMessage {
					t.Errorf("Expected error: %s, got: %s", c.errorMessage, err.Error())
				}
				return
			}

			if cfg.Port != c.want {
				t.Errorf("want %s, got %s", c.want, cfg.Port)
			}
		})
	}
}
