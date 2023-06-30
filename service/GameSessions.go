package service

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/google/uuid"
)

type Session struct {
	ID          string
	Board       [][]map[string]interface{}
	Active      bool
	Turn        int
	Token       string
	Level       int
	Path        []map[string]interface{}
	Destination map[string]interface{}
	Set         []map[string]interface{}
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
		Token:       "TokenCheck00",
		Level:       level,
		Path:        path,
		Destination: des,
		Set:         set,
	}

	sessions[id] = session
	return board, id
}

func generateUniqueToken() string {
	uuid := uuid.New()

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

func TimeOutSessions(id string, turn int) ([][]map[string]interface{}, string, error) {
	session, exists := GetSession(id)
	if !exists {
		return nil, "", errors.New("session not found")
	}

	gameBoard, newToken, newPath, newDes, err := TimeOut(turn, session)
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

func EndSession(id string) {
	delete(sessions, id)
}

func hashToken(o string) string {
	hash := sha256.New()

	hash.Write([]byte(o))

	checksum := hash.Sum(nil)

	checksumStr := hex.EncodeToString(checksum)
	return checksumStr
}
