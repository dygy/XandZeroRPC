package Main

import (
	"net/http"
)

var xorzero = new(Table)

func Main() {
	xorzero.init()
	mux := http.NewServeMux()
	mux.HandleFunc("/api/", Cors)
	http.ListenAndServe(":3333", mux)
}
