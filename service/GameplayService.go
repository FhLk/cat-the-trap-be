package service

import (
	"cat-the-trap-back-end/Algorithm"
	"errors"
	"fmt"
	"math"
)

func UpdateBoard(x, y, turn int, block string, board [][]map[string]interface{}) ([][]map[string]interface{}, string, error) {
	hexagonDisable := []string{"./candy1.svg", "./candy2.svg", "./candy3.svg", "./candy4.svg", "./candy5.svg", "./candy6.svg", "./candy7.svg"}
	if !isInArray(hexagonDisable, block) {
		return nil, "", errors.New("something is wrong")
	}
	board[x][y]["hexagon"] = block
	board[x][y]["block"] = true
	gameBoard := board
	newToken := fmt.Sprintf("TokenCheck0%d", turn)
	if !CatMove(gameBoard) {
		if len(path) == 0 {
			newToken = fmt.Sprintf("TokenCheck0%d", turn-1)
			return gameBoard, newToken, nil
		}
		return gameBoard, "", nil
	}
	return gameBoard, newToken, nil
}

func ResetBoard(level int) [][]map[string]interface{} {
	gameBoard := GameSetup(level)
	return gameBoard
}

func CatMove(board [][]map[string]interface{}) bool {
	currentMove := board[path[0]["x"].(int)][path[0]["y"].(int)]
	path = Algorithm.AStar(currentMove, end, board)
	if len(path) == 0 {
		end = CloseCat(currentMove, board)
		path = Algorithm.AStar(currentMove, end, board)
	} else if len(path) < 5 {
		end = CloseCat(currentMove, board)
		path = Algorithm.AStar(currentMove, end, board)
	} else if len(path) > 7 {
		end = CloseCat(currentMove, board)
		path = Algorithm.AStar(currentMove, end, board)
	}

	if len(path) != 0 {
		previousMove := board[path[0]["x"].(int)][path[0]["y"].(int)]
		previousMove["hexagon"] = hexagonNormal
		previousMove["cat"] = false
		previousMove["block"] = false
		path = path[1:]
		nextMove := board[path[0]["x"].(int)][path[0]["y"].(int)]
		nextMove["cat"] = true
		if CheckLoseGame(nextMove) {
			return false
		}
		return true
	}
	return false
}

func CloseCat(currentCat map[string]interface{}, board [][]map[string]interface{}) map[string]interface{} {
	var filtered []map[string]interface{}
	for _, n := range setDestination {
		if block, ok := n["block"].(bool); !ok || !block {
			filtered = append(filtered, n)
		}
	}
	setDestination = filtered
	distance := math.Inf(1)
	newDestination := end
	for _, n := range setDestination {
		newPath := Algorithm.AStar(currentCat, n, board)
		if len(newPath) != 0 {
			newPath = newPath[1:]
		}
		block, ok := n["block"].(bool)
		newPathLength := len(newPath)
		if newPathLength > 0 &&
			float64(newPathLength) < distance &&
			(!ok || !block) {
			newDestination = newPath[len(newPath)-1]
			distance = float64(newPathLength)
		}
	}
	fmt.Println(newDestination)
	return newDestination
}

func isInArray(arr []string, target string) bool {
	for _, val := range arr {
		if val == target {
			return true
		}
	}
	return false
}

func CheckLoseGame(current map[string]interface{}) bool {
	for _, value := range setDestination {
		if current["x"].(int) == value["x"].(int) && current["y"].(int) == value["y"].(int) {
			return true
		}
	}
	return false
}
