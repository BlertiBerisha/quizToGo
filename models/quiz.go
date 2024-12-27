package models

import (
	"encoding/json"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Command constants
const (
	CommandStart     = "1"
	CommandHighscore = "2"
	CommandTimer     = "3"
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

// GetRandomQuestionsByDifficulty returns random questions filtered by difficulty
func GetRandomQuestionsByDifficulty(difficulty string, questionCount int) []Question {
	// Filter questions by difficulty
	var filteredQuestions []Question
	for _, question := range questions {
		if question.Difficulty == difficulty {
			filteredQuestions = append(filteredQuestions, question)
		}
	}

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// If there are not enough questions of the selected difficulty, return as many as possible
	if len(filteredQuestions) < questionCount {
		questionCount = len(filteredQuestions)
	}

	// Randomly shuffle the questions
	rand.Shuffle(len(filteredQuestions), func(i, j int) {
		filteredQuestions[i], filteredQuestions[j] = filteredQuestions[j], filteredQuestions[i]
	})

	// Return the selected number of random questions
	return filteredQuestions[:questionCount]
}

// GetQuestionsByTopicAndDifficulty returns questions filtered by topic and difficulty
func GetQuestionsByTopicAndDifficulty(topic, difficulty string, count int) []Question {
	var filteredQuestions []Question

	// Filter questions by topic and difficulty
	for _, question := range questions {
		if strings.ToLower(question.Topic) == strings.ToLower(topic) && strings.ToLower(question.Difficulty) == strings.ToLower(difficulty) {
			filteredQuestions = append(filteredQuestions, question)
		}
	}

	// Return the requested number of questions, or all if less than count
	if count <= 0 {
		return filteredQuestions
	}
	if count > len(filteredQuestions) {
		count = len(filteredQuestions)
	}
	return filteredQuestions[:count]
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
