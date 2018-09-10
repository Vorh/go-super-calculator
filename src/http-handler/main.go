package main

import (
	"net/http"
	"go-super-calculator/src/utils"
	"src/github.com/op/go-logging"
	"encoding/json"
)

var log *logging.Logger

func main() {

	http.HandleFunc("/calculate", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func init()  {
	utils.InitLogger()
	log = utils.GetLogger()
}

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	get := r.PostForm.Get("ex")

	hash := make(map[string]string)
	hash["result"] = "ok"
	bytes, _ := json.Marshal(hash)
	w.Header().Set("Content-Type","application/json")
	w.Write(bytes)

	log.Debugf("Handle: %s",get)
}
