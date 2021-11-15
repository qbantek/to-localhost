package routes_test

import (
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/qbantek/to-localhost/internal/routes"
)

func TestIndex(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := routes.NewEngine()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Want %d, got %v", http.StatusOK, w.Code)
	}
}

func TestRedirect(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := routes.NewEngine()

	cases := []struct {
		name           string
		port           string
		httpStatusCode int
	}{
		{"Valid port value", "3000", http.StatusMovedPermanently},
		{"Invalid port value", "invalid", http.StatusBadRequest},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/"+c.port, nil)
			router.ServeHTTP(w, req)

			if c.httpStatusCode != w.Code {
				t.Errorf("Want %v, got %v", c.httpStatusCode, w.Code)
			}
		})
	}
}

func randomInteger(min, max int) int {
	return rand.Intn(max-min) + min
}
