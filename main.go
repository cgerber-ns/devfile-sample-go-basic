package main

import (
	"fmt"
	"net/http"
	"flag"
)
var port = flag.Int("p", 8080, "server port")

func main() {
	flag.Parse()
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/ready", Readied)
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", *port), nil)
}

func Readied(w http.ResponseWriter, r *http.Request) {
	w.WriteResponse(http.StatusOK)
}
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
