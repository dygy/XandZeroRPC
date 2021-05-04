package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Request struct {
	request string
	id uuid.UUID
	action string
}

func parseHeader(r *http.Request) *Request {
	var request = new(Request)
	request.request = r.Header.Get("request")
	request.id, _ = uuid.FromString(r.Header.Get("id"))
	log.Println(request.id, r.Header.Get("id"))

	request.action = r.Header.Get("action")
	return request
}

func RPC(w http.ResponseWriter, req *Request, xorzero *Table) {
	if xorzero.winner != uuid.Nil {
		GETTable(w, xorzero)
		return
	}

	switch req.request {
	case "GETTable": GETTable(w, xorzero)
	case "placeUnit": {parsePlace(req, xorzero);GETTable(w, xorzero)}
	case "refresh": refresh(w, xorzero)
	case "giveSlot": giveSlot(w, xorzero)
	default:
		w.Write([]byte("hello"))
	}
}

func GETTable(w http.ResponseWriter, xorzero *Table) {
	w.Write([]byte(fmt.Sprintf(
		"{ winner: %d, matrix: %d, lastMover: %d }",
		xorzero.winner, xorzero.matrix, xorzero.lastMover,
		)))
}

func refresh(w http.ResponseWriter, xorzero *Table) {
	if xorzero.winner != uuid.Nil {
		xorzero.init()
		GETTable(w, xorzero)
	}
}

func giveSlot(w http.ResponseWriter, xorzero *Table) {
	var slot = xorzero.giveSlot().String()
	w.Write([]byte(slot))
}

func parsePlace(req *Request, xorzero *Table) {
	var array = strings.Split(req.action, "")
	column, _ := strconv.ParseInt(array[1], 10 ,64)
	row, _ := strconv.ParseInt(array[0], 10, 64)
	xorzero.placeUnit(
		row,
		column,
		req.id,
		)
	xorzero.checkWinner()
}
