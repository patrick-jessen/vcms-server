package start

import (
	"log"
	"os/exec"
	"runtime"
)

var browserArgs = []string{
	//"--app=http://localhost:1337/",
	"http://localhost:1337/",
}

// Browser starts the browser.
// It panics if browser cannot be started
func Browser() {
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
