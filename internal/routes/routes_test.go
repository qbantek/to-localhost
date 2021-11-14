package routes_test

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/qbantek/to-localhost/internal/routes"
)

func TestIndex(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := routes.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Want %d, got %v", http.StatusOK, w.Code)
	}
}

func TestRedirect(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := routes.SetupRouter()

	// get a random Port number between 1 and 65535
	rand.Seed(time.Now().UnixNano())
	n := fmt.Sprintf("%d", randomInteger(1, 65535))

	cases := []struct {
		name           string
		port           string
		httpStatusCode int
		locationHeader string
	}{
		{"Default HTTP Port", "80", http.StatusMovedPermanently, "http://localhost"},
		{"A valid Port", n, http.StatusMovedPermanently, "http://localhost:" + n},
		{"Non numeric Port value", "invalid", http.StatusBadRequest, ""},
		{"Port number < 1", "0", http.StatusBadRequest, ""},
		{"Port number > 65535", "65536", http.StatusBadRequest, ""},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/"+c.port, nil)
			router.ServeHTTP(w, req)

			if c.httpStatusCode != w.Code {
				t.Errorf("Want %v, got %v", c.httpStatusCode, w.Code)
			}

			if w.Header().Get("Location") != c.locationHeader {
				t.Errorf("Want %v, got %v", c.locationHeader, w.Header().Get("Location"))
			}
		})
	}
}

func randomInteger(min, max int) int {
	return rand.Intn(max-min) + min
}
