package service

import (
	"cat-the-trap-back-end/Algorithm"
	"fmt"
)

func UpdateBoard(x, y, turn int, block string, board [][]map[string]interface{}) ([][]map[string]interface{}, string) {
	hexagonDisable := []string{"./candy1.svg", "./candy2.svg", "./candy3.svg", "./candy4.svg", "./candy5.svg", "./candy6.svg", "./candy7.svg"}
	fmt.Println(isInArray(hexagonDisable, block))
	if !isInArray(hexagonDisable, block) {
		return nil, ""
	}
	board[x][y]["hexagon"] = block
	board[x][y]["block"] = true
	gameBoard := board
	newToken := fmt.Sprintf("TokenCheck0%d", turn)
	CatMove(gameBoard)
	return gameBoard, newToken
}

func ResetBoard() [][]map[string]interface{} {
	gameBoard := GameSetup()
	return gameBoard
}

func CatMove(board [][]map[string]interface{}) {
	currentMove := board[path[0]["x"].(int)][path[0]["y"].(int)]
	path = Algorithm.AStar(currentMove, end, board)
	previousMove := board[path[0]["x"].(int)][path[0]["y"].(int)]
	previousMove["hexagon"] = hexagonNormal
	previousMove["cat"] = false
	previousMove["block"] = false
	path = path[1:]
	nextMove := board[path[0]["x"].(int)][path[0]["y"].(int)]
	nextMove["cat"] = true
}

func isInArray(arr []string, target string) bool {
	for _, val := range arr {
		if val == target {
			return true
		}
	}
	return false
}
