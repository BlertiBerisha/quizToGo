package controllers

import (
	"fmt"
	"quizToGo/models"
	"quizToGo/views/console"
	"time"
)

// Run does the running of the console application
func Run() {
	err := models.Initialize()
	checkAndHandleError(err)

	// Show the menu when the application starts
	console.Clear()
	console.ShowMenu()

	// Loop to continuously prompt the user for commands
	for {
		executeCommand()
	}
}

func checkAndHandleError(err error) {
	if err != nil {
		console.ShowError(err)
	}
}

func executeCommand() {
	// Ask for the user's input and parse the command
	command := console.AskForInput()
	parseAndExecuteCommand(command)
}

// timeoutChan is used to signal when the timer expires
var timeoutChan chan bool

// answerChan is used to receive the user's answer
var answerChan chan string

// Initialize channels
func initializeChannels() {
	timeoutChan = make(chan bool)
	answerChan = make(chan string)
}

func getAnswerWithTimeout() (string, bool) {
	answerCh := make(chan string, 1)

	go func() {
		answer := console.AskForInput()
		answerCh <- answer
	}()

	select {
	case answer := <-answerCh:
		return answer, false
	case <-time.After(time.Duration(models.TimerValue) * time.Second):
		return "", true
	}
}

func parseAndExecuteCommand(input string) {
	switch {
	case input == models.CommandStart:
		console.Clear()

		// Ask for the player Name
		console.ShowMessage("Please enter your name:")
		PlayerName := console.AskForInput()

		// Ask for and handle topic selection
		console.ShowAvailableTopics()
		topicChoice := console.AskForInput()

		// Ask for and handle difficulty selection
		console.ShowAvailableDifficulties()
		difficultyChoice := console.AskForInput()

		// Ask for the number of questions
		console.ShowHowManyQuestions()
		input := console.AskForInput()
		questionCount := models.StringToInt(input)

		var selectedQuestions []models.Question

		if topicChoice == "Random" {
			selectedQuestions = models.GetRandomQuestionsByDifficulty(difficultyChoice, questionCount)
		} else {
			selectedQuestions = models.GetQuestionsByTopicAndDifficulty(topicChoice, difficultyChoice, questionCount)
		}

		if len(selectedQuestions) > 0 {
			// Show timer information before starting the quiz
			console.Clear()
			console.ShowMessage(fmt.Sprintf("You have %d seconds to answer each question.", models.TimerValue))
			console.ShowMessage("Press Enter to start the quiz...")
			console.AskForInput()

			currentScore := 0

			for _, question := range selectedQuestions {
				console.Clear()
				console.ShowQuestion(question)

				answer, timedOut := getAnswerWithTimeout()

				if timedOut {
					console.ShowMessage("\nTime's up! Question marked as incorrect.")
					console.ShowWrongAnswer(question.Options[question.CorrectIndex])
				} else {
					answerIndex := models.StringToInt(answer) - 1
					if answerIndex == question.CorrectIndex {
						currentScore++
						console.ShowCorrectAnswer()
					} else {
						console.ShowWrongAnswer(question.Options[question.CorrectIndex])
					}
				}

				// Prompt the user to continue after each question
				console.ShowContinue()
				console.AskForInput()
			}

			// Show the final score after the quiz
			console.ShowFinalScore(currentScore, len(selectedQuestions))
			console.ShowContinue()
			console.AskForInput()

			// Add the player's score to the high scores list
			models.AddHighScore(models.Score{
				PlayerName:     PlayerName,
				Score:          currentScore,
				TotalQuestions: len(selectedQuestions),
			})

		} else {
			console.ShowMessage("No questions available.")
		}

		// Return to the main menu
		console.Clear()
		console.ShowMenu()

	case input == models.CommandHighscore:
		// Show high scores
		console.Clear()
		scores := models.GetHighScores()
		console.ShowHighScores(scores)
		console.ShowContinue()
		console.AskForInput()
		console.Clear()
		console.ShowMenu()

	case input == models.CommandTimer:
		// Set default Timer
		// Change timer value
		console.Clear()
		console.ShowMessage("Current timer value (seconds): " + fmt.Sprintf("%d", models.TimerValue))
		console.ShowMessage("Enter new timer value (in seconds):")
		newTimerValue := console.AskForInput()
		newTimerInt := models.StringToInt(newTimerValue)

		// Update the timer value
		models.SetTimerValue(newTimerInt)

		console.ShowMessage("Timer value updated successfully!")
		console.ShowContinue()
		console.AskForInput()
		console.Clear()
		console.ShowMenu()

	case input == models.CommandClear:
		// Clear the view and return to the menu
		console.Clear()
		console.ShowMenu()

	case input == models.CommandQuit:
		// Quit the application
		console.Clear()
		console.ShowGoodbye()
		console.ShutDownNormal()

	default:
		// Handle undefined commands
		console.ShowMessage("Command not defined. Check menu for available commands.")
	}
}
