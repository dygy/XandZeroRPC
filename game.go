package Main

import (
	"log"
)

type Table struct {
	matrix    [][]string
	winner    string
	lastMover uuid.UUID
	players   [2]uuid.UUID
}

func emptyLine() []string {
	var emptyLine []string
	emptyLine = append(emptyLine, "0", "0", "0")
	return emptyLine
}

func (xorzero *Table) init() {
	xorzero.players[0] = uuid.Nil
	xorzero.players[1] = uuid.Nil
	xorzero.lastMover = uuid.Nil
	xorzero.winner = "0"
	xorzero.matrix = xorzero.matrix[:0]
	xorzero.matrix = append(xorzero.matrix, emptyLine(), emptyLine(), emptyLine())
}

func (xorzero *Table) placeUnit(row int64, column int64, playerUuid uuid.UUID) {
	if xorzero.winner != "0" {
		return
	}
	var length = int64(len(xorzero.matrix))
	var isInRange = (column <= length && row <= length) && (column >= 0 && row >= 0)
	var isAllInLobby = !uuid.Equal(uuid.Nil, xorzero.players[0]) && !uuid.Equal(uuid.Nil, xorzero.players[1])
	var isNormalPlayer = xorzero.checkPlayer(playerUuid)
	var isPlayer = xorzero.checkLobby(playerUuid)

	log.Println(isNormalPlayer, playerUuid, xorzero.lastMover)
	if isNormalPlayer && isInRange && isPlayer && isAllInLobby {
		var isEmpty = xorzero.matrix[column][row] == "0"

		if isEmpty {
			xorzero.matrix[column][row] = xorzero.getPlayer(playerUuid)
			xorzero.lastMover = playerUuid
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

func (xorzero *Table) checkLobby(from uuid.UUID) bool {
	return uuid.Equal(xorzero.players[0], from) || uuid.Equal(xorzero.players[1], from)
}

func (xorzero *Table) checkPlayer(from uuid.UUID) bool {
	return !uuid.Equal(xorzero.lastMover, from)
}

func (xorzero *Table) checkRule(e1 string, e2 string, e3 string) {
	if e1 == e2 && e2 == e3 && e1 != "0" {
		xorzero.winner = e1
	}
}

func (xorzero *Table) getPlayer(from uuid.UUID) string {
	if uuid.Equal(xorzero.players[0], from) {
		return "1"
	}
	if uuid.Equal(xorzero.players[1], from) {
		return "2"
	}
	return "0"
}

func (xorzero *Table) giveSlot() uuid.UUID {
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
