package api

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"strings"

	"github.com/gin-gonic/gin"
)

// ConnectHost sets the host cookie on the client and proxies to it.
func ConnectHost(c *gin.Context) {
	fmt.Println("stting cookiesssss")

	protocol := c.Param("protocol")
	host := c.Param("host")
	port := c.Param("port")

	// Sanitize protocol
	switch protocol {
	case "http":
	case "https":
	default:
		log.Panic("Invalid protocol", protocol)
	}

	// Construct URL string
	urlStr := fmt.Sprintf("%s://%s", protocol, host)
	if len(port) > 0 {
		urlStr += ":" + strings.TrimPrefix(port, "/")
	}

	// Parse URL to verify validity
	url, err := url.Parse(urlStr)
	if err != nil {
		panic(err)
	}

	// Check if host is running
	_, err = http.Get(url.String())
	if err != nil {
		c.AbortWithError(http.StatusRequestTimeout, err)
		return
	}

	// Set cookie
	c.SetCookie("vcms", url.String(), 0, "", "", false, true)
	c.Status(http.StatusOK)
}
