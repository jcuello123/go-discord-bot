package battleship

import (
	"go-discord-bot/emoji"
	"testing"
)

func TestCreateNewBoard(t *testing.T) {
	createNewBoard()
	var emptyBoard board

	if gameBoard == emptyBoard{
		t.Errorf("Expected board not to be nil.")
	}

	for i := range gameBoard.cells {
		for j := range gameBoard.cells[i] {
			cellEmoji := gameBoard.cells[i][j].emoji
			empty := gameBoard.cells[i][j].empty
			hit := gameBoard.cells[i][j].hit

			if  cellEmoji != emoji.BLUE_SQUARE {
				t.Errorf("Expected every cell to contain the BLUE_SQUARE emoji, but received: %s", cellEmoji)
			}
			if empty != true {
				t.Errorf("Expected every cell to be empty but received empty: %v", empty)
			}
			if hit != false {
				t.Errorf("Expected every cell to not be hit but received hit: %v", hit)
			}
		}
	}
}

func TestGetBoardAsString(t *testing.T) {
	expectedBoard := ":blue_square:     :blue_square:     :blue_square:     :blue_square:     :blue_square:\n\n" + 
	":blue_square:     :blue_square:     :blue_square:     :blue_square:     :blue_square:\n\n" + 
	":blue_square:     :blue_square:     :blue_square:     :blue_square:     :blue_square:\n\n" +
	":blue_square:     :blue_square:     :blue_square:     :blue_square:     :blue_square:\n\n" +
	":blue_square:     :blue_square:     :blue_square:     :blue_square:     :blue_square:"

	createNewBoard()

	actualBoard := GetBoardAsString()
	if  actualBoard != expectedBoard{
		t.Errorf("Expected %s but received %s", expectedBoard, actualBoard)
	}
}