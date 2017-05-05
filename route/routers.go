package route

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"strings"

	"os"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Assets(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	path := fmt.Sprintf("./assets/%s", params["file"])

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	w.Header().Set("Content-Type", mime.TypeByExtension(params["file"]))
	_, err = io.Copy(w, file)
	if err != nil {
		panic(err)
	}
}

func AssetsClient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	path := fmt.Sprintf("./assets/client/%s", params["file"])

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var ctype = ""
	var split = strings.Split(params["file"], ".")
	var ext = split[len(split)-1]

	switch ext {
	case "css":
		ctype = "text/css"
	case "png":
		ctype = "image/png"
	}

	w.Header().Set("Content-Type", ctype)
	_, err = io.Copy(w, file)
	if err != nil {
		panic(err)
	}
}

var routes = Routes{

	Route{
		"Stop server",
		"GET",
		"/stop",
		Stop,
	},

	Route{
		"Prevent server from stopping",
		"GET",
		"/preventStop",
		PreventStop,
	},

	Route{
		"Assets",
		"GET",
		"/assets/{file}",
		Assets,
	},
	Route{
		"Assets",
		"GET",
		"/assets/client/{file}",
		AssetsClient,
	},

	Route{
		"RootPost",
		"POST",
		"/",
		RootPost,
	},
}
