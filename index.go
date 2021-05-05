package main

import (
	"net/http"
)

var xorzero = new(Table)

func main() {
	xorzero.init()
	mux := http.NewServeMux()
	mux.HandleFunc("/api/", Cors)
	http.ListenAndServe(":3333", mux)
}
