package route

import (
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/gin-gonic/gin"

	"net/url"

	"fmt"

	"github.com/patrick-jessen/vcms-server/api"
)

var r http.Handler

// CreateRouter creates and returns the http router.
func CreateRouter() http.Handler {
	r := gin.Default()

	// API routes
	createAPIRoutes(r.Group("/api"))

	// Reverse proxy
	r.Use(reverseProxy)

	return r
}

func createAPIRoutes(r *gin.RouterGroup) {
	r.GET("/stop", api.Stop)
	r.GET("/preventStop", api.PreventStop)
	r.GET("/connect/:protocol/:host/*port", api.ConnectHost)
}

func reverseProxy(c *gin.Context) {

	noCache(c)

	// Get cookie with host
	appHost, err := c.Cookie("vcms")

	// Check for cookie reset query
	if len(c.Query("vcms")) > 0 {
		c.SetCookie("vcms", "", 0, "", "", false, true)
		appHost = ""
	}

	// Check for missing host cookie
	if err != nil || len(appHost) == 0 {
		// Serve client, skip proxying
		file := fmt.Sprintf("./assets/client%s", c.Request.URL.Path)
		http.ServeFile(c.Writer, c.Request, file)
		c.Header("Cache-Control", "no-cache")
		return
	}

	// Validate URL
	urlStr := appHost //+ c.Request.URL.String()
	url, err := url.Parse(urlStr)
	if err != nil {
		panic(err)
	}

	// Perform proxy
	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.ServeHTTP(c.Writer, c.Request)

	c.Header("Cache-Control", "no-cache")
}

var epoch = time.Unix(0, 0).Format(time.RFC1123)
var noCacheHeaders = map[string]string{
	"Expires":         epoch,
	"Cache-Control":   "no-cache, private, max-age=0",
	"Pragma":          "no-cache",
	"X-Accel-Expires": "0",
}
var etagHeaders = []string{
	"ETag",
	"If-Modified-Since",
	"If-Match",
	"If-None-Match",
	"If-Range",
	"If-Unmodified-Since",
}

func noCache(c *gin.Context) {
	for _, v := range etagHeaders {
		if c.Request.Header.Get(v) != "" {
			c.Request.Header.Del(v)
		}
	}

	for k, v := range noCacheHeaders {
		c.Writer.Header().Set(k, v)
	}
}
