package main

import (
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Request struct {
	request string
	id      uuid.UUID
	action  string
}

func parseHeader(r *http.Request) *Request {
	var request = new(Request)
	request.request = r.Header.Get("request")
	request.id, _ = uuid.FromString(r.Header.Get("id"))
	request.action = r.Header.Get("action")

	return request
}

func RPC(w http.ResponseWriter, req *Request, xorzero *Table) {
	if xorzero.winner != "0" && req.request != "refresh" {
		GETTable(w, xorzero)
		return
	}

	switch req.request {
	case "GETTable":
		GETTable(w, xorzero)
	case "placeUnit":
		parsePlace(w, req, xorzero)
	case "refresh":
		refresh(xorzero, req)
	case "giveSlot":
		giveSlot(w, xorzero)
	default:
		w.Write([]byte(`{"hello": "hello"}`))
	}
}
func matrixToString(matrix [][]string) string {
	var string = "["

	for i := 0; i < len(matrix); i++ {
		string += "[" + strings.Join(matrix[i], ",") + "]"
		if i < len(matrix)-1 {
			string += ","
		}
	}

	return string + "]"
}

func GETTable(w http.ResponseWriter, xorzero *Table) {
	result, _ := json.Marshal(
		"{\"winner\":" + xorzero.winner + ",\"matrix\":" + matrixToString(xorzero.matrix) + "}")
	io.WriteString(w, string(result))
}

func refresh(xorzero *Table, req *Request) {
	if uuid.Equal(req.id, xorzero.players[0]) || uuid.Equal(req.id, xorzero.players[1]) {
		xorzero.init()
	}
}

func giveSlot(w http.ResponseWriter, xorzero *Table) {
	var slot = xorzero.giveSlot().String()
	w.Write([]byte(slot))
}

func parsePlace(w http.ResponseWriter, req *Request, xorzero *Table) {
	var array = strings.Split(req.action, "")
	column, _ := strconv.ParseInt(array[1], 10, 64)
	row, _ := strconv.ParseInt(array[0], 10, 64)
	xorzero.placeUnit(
		row,
		column,
		req.id,
	)
	xorzero.checkWinner()
	GETTable(w, xorzero)
}

func Cors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=ascii")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers, request, action, id")
	var req = parseHeader(r)
	log.Println(req.action)

	RPC(w, req, xorzero)
}
