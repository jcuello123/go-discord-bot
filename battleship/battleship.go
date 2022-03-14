package battleship

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"go-discord-bot/emoji"
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
	"Target a cell by entering **!bs shoot x y** where **x** and **y** are the coordinates."

	gameBoard board 
	started bool = false
	attempts int = 0
	boatsDestroyed int = 0
)

const (
	maxAttempts int = 10 
	totalBoats int = 3
)

func Start() {
	started = true
	createNewBoard()
	spawnBoatsAtRandomLocations()
}

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

// TODO: Add tests for these functions 

func Shoot(args []string) (string, error){
	if !started {
		return "", errors.New("Game hasn't started yet.")
	}

	x, y, err := validateArgs(args)
	if err != nil {
		return "", err
	}

	targetCell := gameBoard.cells[y][x]
	if targetCell.hit {
		return "", errors.New("Already hit.")
	}

	targetCell.hit = true

	attempts++

	if targetCell.empty {
		targetCell.emoji = emoji.X
	} else {
		boatsDestroyed++
		targetCell.emoji = emoji.CRUISE_SHIP
	} 

	gameBoard.cells[y][x] = targetCell


	if targetCell.empty && attempts == maxAttempts {
		reset()
		return "", errors.New("Max attempts reached. Game over.")
	}

	if boatsDestroyed == totalBoats {
		reset()
		return "All boats destroyed! Game over.", nil
	}

	return GetRemainingAttempts(), nil
}

func validateArgs(args []string) (int, int, error){
	if len(args) != 2 {
		return -1, -1, errors.New(fmt.Sprintf("Shoot command takes a total of 2 arguments. Received %d", (len(args))))
	}

	x, err := strconv.Atoi(args[0]) 
	if err != nil {
		return -1, -1, errors.New("Invalid argument for x.") 
	}

	y, err := strconv.Atoi(args[1]) 
	if err != nil {
		return -1, -1, errors.New("Invalid argument for y.")
	}
	
	if x <= 0 || x > len(gameBoard.cells) || y <= 0 || y > len(gameBoard.cells[0]) {
		return -1, -1, errors.New("Coordinates are out of bounds")
	}

	return x - 1, y - 1, nil
}

func reset() {
	attempts = 0
	started = false
}

func getRandInt(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max) 
}

func spawnBoatsAtRandomLocations() {
	spawnedBoats := 0

	for spawnedBoats != totalBoats {
		randX := getRandInt(len(gameBoard.cells))
		randY := getRandInt(len(gameBoard.cells))

		if gameBoard.cells[randX][randY].empty {
			gameBoard.cells[randX][randY].empty = false
			log.Println(randY , randX)
			spawnedBoats++
		}
	}	
}

func GetRemainingAttempts() string{
	remainingAttempts := fmt.Sprintf("Remaining attempts: %d", maxAttempts - attempts)
	return remainingAttempts 
}