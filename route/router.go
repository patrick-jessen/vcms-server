package route

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/patrick-jessen/vcms-server/api"
)

var r http.Handler

// CreateRouter creates and returns the http router.
func CreateRouter() http.Handler {
	r := gin.Default()

	api := r.Group("/api")
	createAPIRoutes(api)
	createRootRoutes(r)

	return r
}

func createAPIRoutes(r *gin.RouterGroup) {
	r.GET("/stop", api.Stop)
	r.GET("/preventStop", api.PreventStop)
}

func createRootRoutes(r *gin.Engine) {
	r.Use(static.Serve("/", static.LocalFile("assets/client", true)))
}
