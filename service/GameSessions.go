package service

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type Session struct {
	ID          string
	Board       [][]map[string]interface{}
	Active      bool
	Turn        int
	TimeOut     bool
	Token       string
	CanPlay     bool
	Level       int
	Path        []map[string]interface{}
	Destination map[string]interface{}
	Set         []map[string]interface{}

	// Add more session-related data as needed
}

var sessions map[string]*Session

func init() {
	sessions = make(map[string]*Session)
}

func StartSession(level int) ([][]map[string]interface{}, string) {
	id := generateUniqueToken()
	board, path, des, set := GameSetup(level)

	session := &Session{
		ID:          id,
		Board:       board,
		Active:      true,
		Turn:        0,
		TimeOut:     false,
		Token:       "TokenCheck00",
		CanPlay:     true,
		Level:       level,
		Path:        path,
		Destination: des,
		Set:         set,
	}

	sessions[id] = session
	return board, id
}

func generateUniqueToken() string {
	// Generate a new UUID
	uuid := uuid.New()

	// Convert the UUID to a string representation
	id := uuid.String()

	return id
}

func GetSession(id string) (*Session, bool) {
	session, exists := sessions[id]
	return session, exists
}

func UpdateBoardSessions(id string, x, y, turn int, block string) ([][]map[string]interface{}, string, error) {
	session, exists := GetSession(id)
	if !exists {
		return nil, "", errors.New("session not found")
	}

	fmt.Println(session.Path)
	gameBoard, newToken, newPath, newDes, err := UpdateBoard(x, y, turn, block, session)
	if err != nil {
		return nil, "", err
	}

	session.Board = gameBoard
	session.Path = newPath
	session.Destination = newDes
	session.Turn = turn
	session.Token = newToken

	return gameBoard, newToken, nil
}

//func TimeOutSessions(id string, turn int) ([][]map[string]interface{}, string, error) {
//	session, exists := GetSession(id)
//	if !exists {
//		return nil, "", errors.New("session not found")
//	}
//
//	gameBoard, newToken, err := TimeOut(turn, session.Board)
//	if err != nil {
//		return nil, "", err
//	}
//
//	session.Board = gameBoard
//
//	return gameBoard, newToken, nil
//}
//
//func ResetBoardSessions(id string, level int) [][]map[string]interface{} {
//	session, exists := GetSession(id)
//	if !exists {
//		return nil
//	}
//
//	gameBoard := ResetBoard(level)
//	session.Board = gameBoard
//
//	return gameBoard
//}

func EndSession(id string) {
	delete(sessions, id)
}
