package battleship

import (
	"go-discord-bot/emoji"
	"strings"
)

type board struct {
	cells [5][5]cell
}

type cell struct {
	empty bool
	hit bool
	emoji string
} 

var (
	 Help = "Start a game by entering **!bs start** \n" +
	"Target a cell by entering **!bs shoot x,y** where **x** and **y** are the coordinates."

	gameBoard board 
)

func init() {
	createNewBoard()
}

// TODO: Add tests for these functions 

func createNewBoard() {
	for i := range gameBoard.cells {
		for j := range gameBoard.cells[i] {
			gameBoard.cells[i][j].empty = true
			gameBoard.cells[i][j].hit = false 
			gameBoard.cells[i][j].emoji = emoji.BLUE_SQUARE
		}
	}
}

func GetBoardAsString() string {
	var result strings.Builder

	for i := range gameBoard.cells {
		for j := range gameBoard.cells[i] {
			c := gameBoard.cells[i][j]
			result.WriteString(c.emoji)
			if j != len(gameBoard.cells) - 1 {
				result.WriteString(" ")
				result.WriteString(" ")
				result.WriteString(" ")
			}
		}
		if i != len(gameBoard.cells) - 1 {
			result.WriteString("\n")
			result.WriteString("\n")
		}
	}

	return result.String()
}