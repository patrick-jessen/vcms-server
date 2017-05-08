package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

var preventStop = make(chan bool, 1)
var stopping = make(chan bool, 1)

func stopFunc() {
	select {
	case <-preventStop:
		return
	case <-time.After(time.Millisecond * 1000):
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

func Save(w http.ResponseWriter, r *http.Request) {
	const StoreFile = "/home/patrick/Code/js/vcms/src/store/laptops.js"

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")

	r.ParseForm()
	store := r.Form.Get("store")
	fmt.Println("body is", store)

	var out bytes.Buffer
	json.Indent(&out, []byte(store), "", "  ")

	data := []byte(`export default ` + string(out.Bytes()))

	ioutil.WriteFile(StoreFile, data, 0644)

	w.WriteHeader(http.StatusOK)
}
