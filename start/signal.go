package start

import (
	"os"
	"os/exec"
	"os/signal"
)

// Signal sets up capture of SIGINT
func Signal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		exec.Command("wmctrl", "-F", "-c", "VCMS").Output()
		os.Exit(0)
	}()
}
