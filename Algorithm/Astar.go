package Algorithm

import (
	"math"
)

type Node struct {
	x, y int
}

func getNeighbors(node Node, gameBoard [][]map[string]interface{}) []map[string]interface{} {
	x := node.x
	y := node.y
	neighbors := make([]map[string]interface{}, 0)
	addNeighbor := func(x, y int) {
		if x >= 0 && x < len(gameBoard) && y >= 0 && y < len(gameBoard[x]) {
			n := gameBoard[x][y]
			if n != nil && !n["block"].(bool) {
				neighbors = append(neighbors, n)
			}
		}
	}
	// Add 6 directions
	addNeighbor(x-1, y)
	addNeighbor(x+1, y)
	addNeighbor(x, y-1)
	addNeighbor(x, y+1)
	if x%2 == 0 {
		addNeighbor(x-1, y-1)
		addNeighbor(x+1, y-1)
	} else {
		addNeighbor(x-1, y+1)
		addNeighbor(x+1, y+1)
	}
	return neighbors
}

func heuristic(a, b Node) int {
	return int(math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y)))
}

func AStar(startNode, endNode map[string]interface{}, gameBoard [][]map[string]interface{}) []map[string]interface{} {
	start := Node{
		x: startNode["x"].(int),
		y: startNode["y"].(int),
	}

	end := Node{
		x: endNode["x"].(int),
		y: endNode["y"].(int),
	}
	var openSet = []Node{start}
	cameFrom := map[Node]Node{}
	gScore := map[Node]int{}
	fScore := map[Node]int{}
	gScore[start] = 0
	fScore[start] = heuristic(start, end)
	for len(openSet) > 0 {
		var current Node
		currentScore := math.MaxInt32
		for _, node := range openSet {
			score := fScore[node]
			if score < currentScore {
				current = node
				currentScore = score
			}
		}
		if current == end {
			p := []Node{end}

			for pathNode := end; pathNode != start; pathNode = cameFrom[pathNode] {
				p = append([]Node{cameFrom[pathNode]}, p...)
			}
			path := make([]map[string]interface{}, len(p))
			for i := range path {
				path[i] = make(map[string]interface{})
				path[i]["x"] = p[i].x
				path[i]["y"] = p[i].y
			}
			return path
		}
		for i, node := range openSet {
			if node == current {
				openSet = append(openSet[:i], openSet[i+1:]...)
				break
			}
		}
		for _, neighbor := range getNeighbors(current, gameBoard) {
			tentativeGScore := gScore[current] + 1
			nb := Node{
				x: neighbor["x"].(int),
				y: neighbor["y"].(int),
			}

			if _, ok := gScore[nb]; tentativeGScore >= gScore[nb] && ok {
				continue
			}
			cameFrom[nb] = current
			gScore[nb] = tentativeGScore
			fScore[nb] = tentativeGScore + heuristic(nb, end)
			openSet = append(openSet, nb)
		}
	}
	return []map[string]interface{}{}
}
