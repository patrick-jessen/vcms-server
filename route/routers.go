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

func serveClientAsset(w http.ResponseWriter, r *http.Request) {
	serveFile(w, r, "./assets/client/")
}

func serveFile(w http.ResponseWriter, r *http.Request, dir string) {
	params := mux.Vars(r)
	relPath := params["path"]
	if len(relPath) == 0 {
		relPath = "index.html"
	}
	path := fmt.Sprintf(dir+"%s", relPath)

	fmt.Println("PATH", path)

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("DOES NOT EXIST")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	defer file.Close()

	pathSplit := strings.Split(path, ".")
	ext := pathSplit[len(pathSplit)-1]
	ext = "." + ext

	mime := mime.TypeByExtension(ext)
	w.Header().Set("Content-Type", mime)
	_, err = io.Copy(w, file)
	if err != nil {
		panic(err)
	}
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AAAAAAAA")
	w.WriteHeader(http.StatusNotFound)

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
		"hu",
		"GET",
		"/",
		serveClientAsset,
	},

	Route{
		"Assets",
		"GET",
		"/{path:.*}",
		serveClientAsset,
	},

	Route{
		"RootPost",
		"POST",
		"/",
		RootPost,
	},
}
