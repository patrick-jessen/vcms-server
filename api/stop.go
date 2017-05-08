package api

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

const stopTimeoutMs = 1000

var preventStop = make(chan bool, 1)
var stopping = make(chan bool, 1)

func stopFunc() {
	select {
	case <-preventStop:
		return
	case <-time.After(time.Millisecond * stopTimeoutMs):
		fmt.Println("Exited because browser was closed")
		os.Exit(0)
	}
}

// Stop tells the server to stop after a timeout.
func Stop(c *gin.Context) {
	if len(preventStop) > 0 {
		<-preventStop // Empty channel
	}

	go stopFunc()
	c.Status(200)
}

// PreventStop tells server to not stop afterall.
func PreventStop(c *gin.Context) {
	select {
	case preventStop <- true:
	default:
	}
	c.Status(200)
}
