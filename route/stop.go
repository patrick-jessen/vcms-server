package route

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

var preventStop = make(chan bool, 1)
var stopping = make(chan bool, 1)

func stopFunc() {
	select {
	case <-preventStop:
		return
	case <-time.After(time.Millisecond * 100):
		fmt.Println("Exited because browser was closed")
		os.Exit(0)
	}
}

func Stop(w http.ResponseWriter, r *http.Request) {
	if len(preventStop) > 0 {
		<-preventStop // Empty channel
	}

	go stopFunc()
}

func PreventStop(w http.ResponseWriter, r *http.Request) {
	select {
	case preventStop <- true:
	}
}
