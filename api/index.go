package handler

import (
	"net/http"
)

var xorzero = new(Table)

func Cors(w http.ResponseWriter, r *http.Request) {
	xorzero.init()
	w.Header().Set("Content-Type", "text/html; charset=ascii")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers, request, action, id")
	var req = parseHeader(r)

	rpc(w, req, xorzero)
}
