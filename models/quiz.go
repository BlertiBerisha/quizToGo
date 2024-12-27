package models

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"
)

// Command constants
const (
	CommandStart     = "1"
	CommandHighscore = "2"
	CommandClear     = "c"
	CommandQuit      = "q"
)

// Question represents a single quiz question
type Question struct {
	ID           int      `json:"id"`
	Text         string   `json:"text"`
	Options      []string `json:"options"`
	CorrectIndex int      `json:"correctIndex"`
	Topic        string   `json:"topic"`
	Difficulty   string   `json:"difficulty"`
}

// Score represents a player's quiz score
type Score struct {
	PlayerName     string
	Score          int
	TotalQuestions int
}

// File name for questions storage
const QuestionsFileName = "questions.json"

// Global variables
var (
	questions  []Question
	highScores []Score
)

// Initialize loads questions from file
func Initialize() error {
	var err error
	questions, err = loadQuestionsFromFile()
	if err != nil {
		return err
	}
	return nil
}

func loadQuestionsFromFile() ([]Question, error) {
	file, err := os.ReadFile(QuestionsFileName)
	if err != nil {
		return nil, err
	}

	var loadedQuestions []Question
	err = json.Unmarshal(file, &loadedQuestions)
	if err != nil {
		return nil, err
	}

	return loadedQuestions, nil
}

// GetQuestions returns a slice of questions based on count
func GetQuestions(count int) []Question {
	if count <= 0 {
		return []Question{}
	}
	if count > len(questions) {
		count = len(questions)
	}
	return questions[:count]
}

// GetHighScores returns all saved high scores
func GetHighScores() []Score {
	return highScores
}

// AddHighScore adds a new high score
func AddHighScore(score Score) {
	highScores = append(highScores, score)
}

// StringToInt converts a string to an integer value
func StringToInt(info string) int {
	infoTrimmed := strings.TrimSpace(info)
	aInt, _ := strconv.Atoi(infoTrimmed)
	return aInt
}
