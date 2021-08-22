package main

import (
	"fmt"
	"os"
)

// Due to the game design, the size of the board also is the max possible value, and is also used as such
/*const size = 9
const emptyValue = 0

type game struct {
	grid [size][size]int
}*/

func main() {

	arguments := os.Args

		var newStr string
	
		for i := 1; i < len(arguments); i++ {
			if len(arguments) == 10 {
				for j := 0; j <= 9; j++ {
					for _, char := range arguments[i] {
						if char >= '1' && char <= '9' || char == '.' {
							newStr = string(arguments[i])
						} else {
							fmt.Println(j)
							fmt.Println("Number is duplicated in one argument")
						}
					}
				}
			} else {
				fmt.Println("Invalid number of arguments")
			}
			fmt.Println(newStr)
		}
}
	
	// test "096008510" "096008510" "110008510" "096222210" "033308510" "000008510" "096008510" "475085101" "096425880"
	//loadAndSolveGame(newStr)

/*func loadAndSolveGame(rawData string) {
	game := loadGame(rawData)
	fmt.Println("Initial board:")
	displayGame(game)
	fmt.Println("Solving it...")
	solution := solve(game)
	fmt.Println("Solution:")
	displayGame(solution)
}

func displayGame(game game) {
	for i := 0; i < len(game.grid); i++ {
		fmt.Println(strings.Trim(strings.ReplaceAll(strings.Join(strings.Fields(fmt.Sprint(game.grid[i])), " "), "0", "_"), "[]"))
	}
}


//Takes a string as an input, made of numbers, where each number is a case in the grid (0 means empty)
//E.g.: 004300209005009001070060043006002087190007400050083000600000105003508690042910300

func loadGame(rawData string) game {
	game := game{}
	array := strings.Split(rawData, "")
	if len(array) != (size * size) {
		return game
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			val, err := strconv.Atoi(array[size*i+j])
			if err == nil && val >= emptyValue && val <= size {
				game.grid[i][j] = val
			} else {
				return game
			}
		}
	}
	if !isValidGame(game) {
		return game
	}
	return game
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

func solve(game game) game {
	if isGameOver(game) {
		return game
	}

	i, j := findEmptyCase(game)

	for val := 1; val <= size; val++ {
		game.grid[i][j] = val
		if isValidGame(game) {
			solve(game)
		}
		game.grid[i][j] = emptyValue
	}

	return game
}

func findEmptyCase(game game) (int, int) {
	for i, line := range game.grid {
		for j, val := range line {
			if val == emptyValue {
				return i, j
			}
		}
	}
	return -1, -1
}
*/

