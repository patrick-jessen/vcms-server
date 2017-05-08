package start

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/patrick-jessen/vcms-server/route"
)

const serverPort = 1337

// Server starts the server.
// Panics if server could not be started.
func Server() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", serverPort))
	if err != nil {
		log.Panicln("Could not start server", err)
	}

	log.Println("Server started on port", serverPort)
	err = http.Serve(listener, route.CreateRouter())
	log.Panicln(err)
}
