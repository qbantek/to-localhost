package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qbantek/to-localhost/internal/localurl"
	"github.com/qbantek/to-localhost/internal/projectpath"
)

// NewEngine returns an Engine (Server) instance with the Logger and Recovery
// middleware already attached and all the routes defined.
func NewEngine() *gin.Engine {
	engine := gin.Default()

	engine.StaticFile("/favicon.ico", projectpath.RootPath+"/static/favicon.ico")
	engine.LoadHTMLGlob(projectpath.RootPath + "/templates/*.tmpl.html")
	engine.Static("/static", projectpath.RootPath+"/static")
	engine.GET("/health", HealthCheck)
	engine.GET("/:port", Redirect)
	engine.GET("/", Index)

	return engine
}

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl.html", nil)
}

func HealthCheck(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

func Redirect(c *gin.Context) {
	url, err := localurl.NewURL(c.Param("port"))
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.tmpl.html", gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusMovedPermanently, url.String())
}
