package controllers

import (
	"fmt"
	"quizToGo/models"
	"quizToGo/views/console"
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

func parseAndExecuteCommand(input string) {
	switch {
	case input == models.CommandStart:
		// Start a new quiz game
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

		// Debugging print
		fmt.Println("Topic:", topicChoice, "Difficulty:", difficultyChoice, "Number of Questions:", questionCount)

		var questions []models.Question

		if topicChoice == "Random" {
			// Get random questions from all topics, respecting difficulty
			questions = models.GetRandomQuestionsByDifficulty(difficultyChoice, questionCount)
		} else {
			// Get questions based on the chosen topic and difficulty
			questions = models.GetQuestionsByTopicAndDifficulty(topicChoice, difficultyChoice, questionCount)
		}

		// Proceed with the quiz
		if len(questions) > 0 {
			currentScore := 0
			for _, question := range questions {
				console.Clear()
				console.ShowQuestion(question)

				// Get the answer from the user
				answer := console.AskForInput()
				answerIndex := models.StringToInt(answer) - 1

				// Check if the answer is correct
				if answerIndex == question.CorrectIndex {
					currentScore++
					console.ShowCorrectAnswer()
				} else {
					console.ShowWrongAnswer(question.Options[question.CorrectIndex])
				}

				// Prompt the user to continue after each question
				console.ShowContinue()
				console.AskForInput()
			}

			// Show the final score after the quiz
			console.ShowFinalScore(currentScore, len(questions))
			console.ShowContinue()
			console.AskForInput()

			// Add the player's score to the high scores list
			models.AddHighScore(models.Score{
				PlayerName:     PlayerName,
				Score:          currentScore,
				TotalQuestions: len(questions),
			})

		} else {
			// Handle case where no questions are available
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
		console.Clear()

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
