package main

import (
	"fmt"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

var logFatal = log.Fatal

// Due to the game design, the size of the board also is the max possible value, and is also used as such
const size = 9
const emptyValue = 0

type game struct {
	grid [size][size]int
}

func main() {

	arguments := os.Args

	var newStr string
	isReapeated := false
	isValidChar := false
        isValidFormat:= true

	if len(arguments) == 1 {
		fmt.Println("Error")
	}

	for i := 1; i < len(arguments); i++ {

		if len(arguments) == 10 && len(arguments[i]) == 9 {

			if !isReapeated && !isValidChar {

				for j := 0; j < len(arguments[i]); j++ {

					if string(arguments[i][j]) == "." || (string(arguments[i][j]) > "0" && string(arguments[i][j]) <= "9") && !isValidChar {

						strReapeated := strings.Count(arguments[i], string(arguments[i][j]))

						if strReapeated > 1 && string(arguments[i][j]) > "0" && string(arguments[i][j]) <= "9" && !isReapeated {
							fmt.Println("Error")
							isReapeated = true
                                                        isValidFormat= false
							break
						}

					} else {
						fmt.Println("Error")
						isValidChar = true
						break
					}
				}
			}

			newStr += string(arguments[i])
		} else {
			fmt.Println("Error")
			break
		}

	}

	if len(newStr) == 81 && isValidFormat {
		strToResolve := strings.ReplaceAll(newStr, ".", "0")
		loadAndSolveGame(strToResolve)

	}
}

func loadAndSolveGame(rawData string) {
	game, err := loadGame(rawData)
	if err != nil {
		logFatal(err)
	}
	solution, err := solve(game)
	if err != nil {
		logFatal(err)
	}
	displayGame(solution)
}

func displayGame(game game) {
	for i := 0; i < len(game.grid); i++ {
		fmt.Println(strings.Trim(strings.ReplaceAll(strings.Join(strings.Fields(fmt.Sprint(game.grid[i])), " "), "0", "_"), "[]"))
	}
}

func loadGame(rawData string) (game, error) {
	game := game{}
	array := strings.Split(rawData, "")
	if len(array) != (size * size) {
		return game, errors.New("incorrect size for game raw data")
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			val, err := strconv.Atoi(array[size*i+j])
			if err == nil && val >= emptyValue && val <= size {
				game.grid[i][j] = val
			} else {
				return game, errors.New("unable to parse valid int")
			}
		}
	}
	if !isValidGame(game) {
		return game, errors.New("invalid game")
	}
	return game, nil
}

func isValidGame(game game) bool {
	for _, line := range game.grid {
		if hasDuplicates(line) {
			return false
		}
	}

	for i := 0; i < size; i++ {
		var column [size]int
		for j := 0; j < size; j++ {
			column[j] = game.grid[j][i]
		}
		if hasDuplicates(column) {
			return false
		}
	}

	return true
}

func hasDuplicates(array [size]int) bool {
	for value := 1; value <= size; value++ {
		found := false
		for _, val := range array {
			if val == value {
				if found {
					return true
				}
				found = true
			}
		}
	}
	return false
}

func isGameOver(game game) bool {
	for _, line := range game.grid {
		for _, val := range line {
			if val == emptyValue {
				return false
			}
		}
	}
	return true
}

func solve(game game) (game, error) {
	if isGameOver(game) {
		return game, nil
	}

	i, j, err := findEmptyCase(game)
	if err != nil {
		return game, err
	}

	for val := 1; val <= size; val++ {
		game.grid[i][j] = val
		if isValidGame(game) {
			solution, err := solve(game)
			if err == nil {
				return solution, nil
			}
		}
		game.grid[i][j] = emptyValue
	}

	return game, errors.New("unable to solve game")
}

func findEmptyCase(game game) (int, int, error) {
	for i, line := range game.grid {
		for j, val := range line {
			if val == emptyValue {
				return i, j, nil
			}
		}
	}
	return -1, -1, errors.New("no empty case found")
}


