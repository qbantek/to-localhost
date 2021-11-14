package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/qbantek/to-localhost/internal/projectpath"
)

const (
	MIN_PORT      = 1
	MAX_PORT      = 65535
	LOCALHOST_URL = "http://localhost"
)

type redirectParams struct {
	Port string
}

type portError struct {
	port string
}

func (e *portError) Error() string {
	return fmt.Sprintf("Invalid port value: %s", e.port)
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob(projectpath.RootPath + "/templates/*.tmpl.html")
	router.Static("/static", projectpath.RootPath+"/static")
	router.GET("/:port", Redirect)
	router.GET("/", Index)

	return router
}

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl.html", nil)
}

func Redirect(c *gin.Context) {
	r := redirectParams{Port: c.Param("port")}

	url, err := r.redirectUrl()
	if err != nil {
		errorPage(c, err)
		return
	}

	c.Redirect(http.StatusMovedPermanently, url)
}

func errorPage(c *gin.Context, err error) {
	c.HTML(http.StatusBadRequest, "error.tmpl.html", gin.H{"error": err.Error()})
}

func (r *redirectParams) redirectUrl() (string, error) {
	url := LOCALHOST_URL

	if r.Port != "" && r.Port != "80" {
		portInt, err := strconv.Atoi(r.Port)
		if err != nil || portInt < MIN_PORT || portInt > MAX_PORT {
			return "", &portError{r.Port}
		}
		url += ":" + r.Port
	}

	return url, nil
}
