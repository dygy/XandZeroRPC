package main

import (
	"github.com/satori/go.uuid"
	"log"
)

type Table struct {
	matrix [3][3]uuid.UUID
	winner uuid.UUID
	lastMover uuid.UUID
	players [2]uuid.UUID
}

func (xorzero *Table) init() {
	xorzero.players[0] = uuid.Nil
	xorzero.players[1] = uuid.Nil
	xorzero.lastMover = uuid.Nil
	xorzero.winner = uuid.Nil
	for i := 0; i < len(xorzero.matrix); i++ {
		for j := 0; j < len(xorzero.matrix); j++ {
			xorzero.matrix[i][j] = uuid.Nil
		}
	}
}

func (xorzero *Table) placeUnit(row int64, column int64, playerIndex uuid.UUID) {
	var length = int64(len(xorzero.matrix))
	var isInRange = (column <= length && row <= length ) && (column >= 0 && row >= 0)
	var isNormalPlayer = !uuid.Equal(playerIndex, xorzero.lastMover)
	var isPlayer = uuid.Equal(playerIndex, xorzero.players[0]) || uuid.Equal(playerIndex, xorzero.players[1])
	log.Println(isNormalPlayer, isInRange, isPlayer, playerIndex)

	if  isNormalPlayer && isInRange && isPlayer {
		var isEmpty = uuid.Equal(xorzero.matrix[column][row], uuid.Nil)

		if isEmpty {
			xorzero.matrix[column][row] = playerIndex
			xorzero.lastMover = playerIndex
		}
	}
}

func (xorzero *Table) checkWinner() {
	for i := 0; i < len(xorzero.matrix); i++ {
		xorzero.checkRule(xorzero.matrix[i][0], xorzero.matrix[i][1], xorzero.matrix[i][2])
		xorzero.checkRule(xorzero.matrix[0][i], xorzero.matrix[1][i], xorzero.matrix[2][i])
	}
	xorzero.checkRule(xorzero.matrix[0][0], xorzero.matrix[1][1], xorzero.matrix[2][2])
	xorzero.checkRule(xorzero.matrix[0][2], xorzero.matrix[1][1], xorzero.matrix[2][0])

}

func (xorzero *Table) checkPlayer(from uuid.UUID) bool {
	return uuid.Equal(xorzero.lastMover, from)
}

func (xorzero *Table) checkRule(e1 uuid.UUID, e2 uuid.UUID, e3 uuid.UUID)  {
	if uuid.Equal(e1, e2) && uuid.Equal(e2, e3) && !uuid.Equal(e1, uuid.Nil) { xorzero.winner = e1 }
}

func (xorzero *Table) giveSlot() uuid.UUID  {
	myuuid, _ := uuid.NewV4()
	if uuid.Equal(xorzero.players[0], uuid.Nil) {
		xorzero.players[0] = myuuid
		return myuuid
	} else if uuid.Equal(xorzero.players[1], uuid.Nil) {
		xorzero.players[1] = myuuid
		xorzero.lastMover = myuuid
		return myuuid
	} else {
		return myuuid
	}
}