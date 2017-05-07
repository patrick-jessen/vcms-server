package main

import (
	"log"
	"net"
	"net/http"
	"runtime"

	"os/exec"

	"fmt"

	sw "github.com/patrick-jessen/vcms-server/route"
)

var browserArgs = []string{"--app=http://localhost:1337/"}

const serverPort = 1337

func startServer() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", serverPort))
	if err != nil {
		log.Panicln("Could not start server", err)
	}
	log.Println("Server started on port", serverPort)
	router := sw.NewRouter()
	err = http.Serve(listener, router)
	log.Panicln(err)
}

func startBrowser() {
	var browserCmds []string

	// Get commands for google chrome
	os := runtime.GOOS
	switch os {
	case "linux":
		browserCmds = []string{"google-chrome", "chromium-browser"}
	case "windows":
		fallthrough
	case "darwin":
		fallthrough
	default:
		log.Panicln("Not supported on platform:", os)
	}

	// Try starting browser
	var started = false
	for _, cmdName := range browserCmds {
		cmd := exec.Command(cmdName, browserArgs...)
		err := cmd.Start()
		if err == nil {
			started = true
			break
		}
	}

	if !started {
		log.Panicln("Could not start browser")
	}
}

func main() {
	startBrowser()
	startServer()
}
