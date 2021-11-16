package config_test

import (
	"os"
	"testing"

	"github.com/qbantek/to-localhost/internal/config"
)

func TestNewConfig(t *testing.T) {
	cases := []struct {
		name string
		env  string
		want string
	}{
		{
			name: "DefaultPort",
			env:  "",
			want: "5000",
		},
		{
			name: "CustomPort",
			env:  "8080",
			want: "8080",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			os.Setenv("PORT", c.env)
			cfg := config.NewConfig()
			if cfg.Port != c.want {
				t.Errorf("want %s, got %s", c.want, cfg.Port)
			}
		})
	}

	cfg := config.NewConfig()
	if cfg == nil {
		t.Error("config.NewConfig() returned nil")
	}
}
