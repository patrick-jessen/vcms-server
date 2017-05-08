package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

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
