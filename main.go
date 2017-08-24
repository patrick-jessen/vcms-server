package main

import (
	"github.com/patrick-jessen/vcms-server/start"
)

func main() {
	start.Signal()
	//start.Browser()
	start.Server()
}
