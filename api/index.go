package handler

import (
	"net/http"
)

var xorzero = new(Table)

func Cors(w http.ResponseWriter, r *http.Request) {
	if !xorzero.isInitialised {
		xorzero.init()
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers, request, action, id")
	rpc(w, parseHeader(r), xorzero)
}
