package service

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
)

func generateBoard() [][]map[string]interface{} {
	board := make([][]map[string]interface{}, 11)
	for i := 0; i < 11; i++ {
		board[i] = make([]map[string]interface{}, 11)
		for j := 0; j < 11; j++ {
			board[i][j] = map[string]interface{}{
				"x":     i,
				"y":     j,
				"block": false,
				"cat":   false,
			}
		}
	}
	return board
}

func divideBoardIntoFour(gameBoard [][]map[string]interface{}) ([][][]map[string]interface{}, error) {
	if len(gameBoard) != 11 || len(gameBoard[0]) != 11 {
		return nil, errors.New("game board must be 11x11")
	}

	Q_TOP := gameBoard[0:5]
	Q_BOTTOM := gameBoard[6:11]

	Q1 := make([][]map[string]interface{}, 5)
	Q2 := make([][]map[string]interface{}, 5)
	Q3 := make([][]map[string]interface{}, 5)
	Q4 := make([][]map[string]interface{}, 5)

	for i := 0; i < len(Q_TOP); i++ {
		Q1[i] = Q_TOP[i][0:5]
		Q2[i] = Q_TOP[i][6:11]
	}

	for i := 0; i < len(Q_BOTTOM); i++ {
		Q3[i] = Q_BOTTOM[i][0:5]
		Q4[i] = Q_BOTTOM[i][6:11]
	}

	return [][][]map[string]interface{}{Q1, Q2, Q3, Q4}, nil
}

func randomBlock(Q [][][]map[string]interface{}, level int) []map[string]interface{} {
	var blocks []map[string]interface{}
	var countBlocks int
	switch level {
	case 1:
		countBlocks = 4
	case 2:
		countBlocks = 3
	case 3:
		countBlocks = 2
	default:
		countBlocks = 1
	}
	for i := 0; i < len(Q); i++ {
		part := Q[i]
		var partBlocks []map[string]interface{}
		for len(partBlocks) < countBlocks {
			row := rand.Intn(len(part))
			col := rand.Intn(len(part[0]))
			block := part[row][col]
			block["block"] = true
			partBlocks = append(partBlocks, block)
		}
		blocks = append(blocks, partBlocks...)
	}

	sort.Slice(blocks, func(i, j int) bool {
		return blocks[i]["x"].(int) < blocks[j]["x"].(int)
	})

	return blocks
}

func Destination(gameBoard [][]map[string]interface{}) []map[string]interface{} {
	BOARD_SIZE := len(gameBoard[0])
	setDestination := make([]map[string]interface{}, 0)
	for i := 0; i < BOARD_SIZE; i++ {
		setDestination = append(setDestination, gameBoard[0][i], gameBoard[i][0], gameBoard[BOARD_SIZE-1][i], gameBoard[i][BOARD_SIZE-1])
	}

	uniqueDestinations := make(map[string]map[string]interface{})
	for _, dest := range setDestination {
		uniqueDestinations[fmt.Sprintf("%v", dest)] = dest
	}

	uniqueSlice := make([]map[string]interface{}, 0, len(uniqueDestinations))
	for _, dest := range uniqueDestinations {
		uniqueSlice = append(uniqueSlice, dest)
	}

	sort.Slice(uniqueSlice, func(i, j int) bool {
		return uniqueSlice[i]["x"].(int) < uniqueSlice[j]["x"].(int)
	})

	destination := make([]map[string]interface{}, 0)
	for _, dest := range uniqueSlice {
		if !dest["block"].(bool) {
			destination = append(destination, dest)
		}
	}

	return destination
}

func GameSetup() ([][]map[string]interface{}, []map[string]interface{}, map[string]interface{}, error) {
	// Generate game board
	gameBoard := generateBoard()

	// Add cat to game board
	gameBoard[5][5]["cat"] = true

	// Divide game board into four quadrants
	Q, err := divideBoardIntoFour(gameBoard)
	if err != nil {
		return nil, nil, nil, nil
	}

	// Generate blocks for quadrants
	block := randomBlock(Q, 1)

	// Set of destinations
	setDestination := Destination(gameBoard)

	// Choose a random destination
	destination := setDestination[rand.Intn(len(setDestination))]

	return gameBoard, block, destination, nil
}