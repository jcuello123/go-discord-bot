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

func TestShoot(t *testing.T) {
	t.Run("Game hasnt started", func(t *testing.T) {
		args := []string{"1", "1"}

		msg, err := Shoot(args)
		if msg != "" {
			t.Errorf("Expected: \"\" got: %s", msg)
		}

		expectedErr := "Game hasn't started yet."
		if err.Error() != expectedErr {
			t.Errorf("Expected: %s Received: %s", expectedErr, err.Error())
		} 
	})

	started = true

	t.Run("Invalid argument for x", func(t *testing.T) {
		args := []string{"a", "b"}

		msg, err := Shoot(args)
		if msg != "" {
			t.Errorf("expected: \"\" got: %s", msg)
		}

		expectedErr := "Invalid argument for x."
		if err.Error() != expectedErr {
			t.Errorf("expected: %s got: %s", expectedErr, err.Error())
		}
	})

	t.Run("Invalid argument for y", func(t *testing.T) {
		args := []string{"1", "b"}

		msg, err := Shoot(args)
		if msg != "" {
			t.Errorf("expected: \"\" got: %s", msg)
		}

		expectedErr := "Invalid argument for y."
		if err.Error() != expectedErr {
			t.Errorf("expected: %s got: %s", expectedErr, err.Error())
		}
	})

	t.Run("Out of bounds x", func(t *testing.T) {
		args := []string{"-4", "3"}

		msg, err := Shoot(args)
		if msg != "" {
			t.Errorf("expected: \"\" got: %s", msg)
		}

		expectedErr :=  "Coordinates are out of bounds"

		if err.Error() != expectedErr {
			t.Errorf("expected: %s got: %s", expectedErr, err.Error())
		}
	})

	t.Run("Out of bounds y", func(t *testing.T) {
		args := []string{"3", "-2"}

		msg, err := Shoot(args)
		if msg != "" {
			t.Errorf("expected: \"\" got: %s", msg)
		}

		expectedErr :=  "Coordinates are out of bounds"

		if err.Error() != expectedErr {
			t.Errorf("expected: %s got: %s", expectedErr, err.Error())
		}
	})

	t.Run("Hit", func(t *testing.T) {
		args := []string{"1", "1"}
		createNewBoard()

		msg, err := Shoot(args)
		if msg == "" {
			t.Errorf("expected: Remaining attempts got: %s", msg)
		}

		if err != nil {
			t.Errorf("expected: %v got: %s", nil, err.Error())
		}

		// if gameBoard.cells[1][1].emoji == emoji.BLUE_SQUARE {
		// 	t.Errorf("expected: %s or %s got: %s", emoji.X, emoji.CRUISE_SHIP, gameBoard.cells[1][1].emoji)
		// }

		// if !gameBoard.cells[1][1].hit {
		// 	t.Errorf("expected: %v got: %v", true, false)
		// }
	})
}