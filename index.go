package main

import (
	"log"
	"net/http"
)

func main() {
	var xorzero = new(Table)
	xorzero.init()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var req = parseHeader(r)
		RPC(w, req, xorzero)
		log.Println(req)
	})
	log.Fatal(http.ListenAndServe(":3333", nil))
}



