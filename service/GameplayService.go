package service

import (
	"cat-the-trap-back-end/Algorithm"
	"errors"
	"fmt"
	"math"
)

func UpdateBoard(x, y, turn int, block string, session *Session) ([][]map[string]interface{}, string, []map[string]interface{}, map[string]interface{}, error) {
	hexagonDisable := []string{"./candy1.svg", "./candy2.svg", "./candy3.svg", "./candy4.svg", "./candy5.svg", "./candy6.svg", "./candy7.svg"}
	if !isInArray(hexagonDisable, block) {
		return nil, "", nil, nil, errors.New("something is wrong")
	}
	board := session.Board
	path := session.Path
	des := session.Destination
	set := session.Set
	board[x][y]["hexagon"] = block
	board[x][y]["block"] = true
	o := fmt.Sprintf("0%d", turn)
	newToken := hashToken("TokenCheck") + o
	move, newPath, newDes := CatMove(board, path, des, set)
	if !move {
		if len(newPath) == 0 {
			o = fmt.Sprintf("0%d", turn-1)
			newToken = hashToken("TokenCheck") + o
			return board, newToken, nil, nil, nil
		}
		return board, "", newPath, newDes, nil
	}
	return board, newToken, newPath, newDes, nil
}

func TimeOut(turn int, session *Session) ([][]map[string]interface{}, string, []map[string]interface{}, map[string]interface{}, error) {
	board := session.Board
	path := session.Path
	des := session.Destination
	set := session.Set
	o := fmt.Sprintf("0%d", turn)
	newToken := hashToken("TokenCheck") + o
	move, newPath, newDes := CatMove(board, path, des, set)
	if !move {
		if len(newPath) == 0 {
			o = fmt.Sprintf("0%d", turn-1)
			newToken = hashToken("TokenCheck") + o
			return board, newToken, nil, nil, nil
		}
		return board, "", newPath, newDes, nil
	}
	return board, newToken, newPath, newDes, nil
}

func CatMove(board [][]map[string]interface{}, path []map[string]interface{}, des map[string]interface{}, set []map[string]interface{}) (bool, []map[string]interface{}, map[string]interface{}) {
	currentMove := board[path[0]["x"].(int)][path[0]["y"].(int)]
	path = Algorithm.AStar(currentMove, des, board)
	if len(path) == 0 {
		des = CloseCat(currentMove, board, set, des)
		path = Algorithm.AStar(currentMove, des, board)
	} else if len(path) < 5 {
		des = CloseCat(currentMove, board, set, des)
		path = Algorithm.AStar(currentMove, des, board)
	} else if len(path) > 7 {
		des = CloseCat(currentMove, board, set, des)
		path = Algorithm.AStar(currentMove, des, board)
	}

	if len(path) != 0 {
		previousMove := board[path[0]["x"].(int)][path[0]["y"].(int)]
		previousMove["hexagon"] = hexagonNormal
		previousMove["cat"] = false
		previousMove["block"] = false
		path = path[1:]
		nextMove := board[path[0]["x"].(int)][path[0]["y"].(int)]
		nextMove["cat"] = true
		if CheckLoseGame(nextMove, set) {
			return false, path, des
		}
		return true, path, des
	}
	return false, nil, nil
}

func CloseCat(currentCat map[string]interface{}, board [][]map[string]interface{}, set []map[string]interface{}, des map[string]interface{}) map[string]interface{} {
	var filtered []map[string]interface{}
	for _, n := range set {
		if block, ok := n["block"].(bool); !ok || !block {
			filtered = append(filtered, n)
		}
	}
	set = filtered
	distance := math.Inf(1)
	newDestination := des
	for _, n := range set {
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

func CheckLoseGame(current map[string]interface{}, set []map[string]interface{}) bool {
	for _, value := range set {
		if current["x"].(int) == value["x"].(int) && current["y"].(int) == value["y"].(int) {
			return true
		}
	}
	return false
}
