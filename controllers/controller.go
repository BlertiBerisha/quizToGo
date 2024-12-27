package controllers

import (
	"quizToGo/models"
	"quizToGo/views/console"
)

// Run does the running of the console application
func Run() {
	err := models.Initialize()
	checkAndHandleError(err)

	console.Clear()
	console.ShowMenu()

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
	command := console.AskForInput()
	parseAndExecuteCommand(command)
}

func parseAndExecuteCommand(input string) {
	switch {
	case input == models.CommandStart:
		// Start a new quiz game
		console.Clear()
		console.ShowQuizSetup()

		// Get number of questions
		input := console.AskForInput()
		questionCount := models.StringToInt(input)

		// Get questions and start quiz
		questions := models.GetQuestions(questionCount)
		currentScore := 0

		for _, question := range questions {
			console.Clear()
			console.ShowQuestion(question)

			answer := console.AskForInput()
			answerIndex := models.StringToInt(answer) - 1

			if answerIndex == question.CorrectIndex {
				currentScore++
				console.ShowCorrectAnswer()
			} else {
				console.ShowWrongAnswer(question.Options[question.CorrectIndex])
			}
			console.ShowContinue()
			console.AskForInput()
		}

		// Show final score
		console.ShowFinalScore(currentScore, len(questions))
		console.ShowContinue()
		console.AskForInput()

		// Return to menu
		console.Clear()
		console.ShowMenu()

	case input == models.CommandHighscore:
		// View highscores
		console.Clear()
		scores := models.GetHighScores()
		console.ShowHighScores(scores)
		console.ShowContinue()
		console.AskForInput()
		console.Clear()
		console.ShowMenu()

	case input == models.CommandClear:
		// Clear view and show menu
		console.Clear()
		console.ShowMenu()

	case input == models.CommandQuit:
		// Quit application
		console.Clear()
		console.ShowGoodbye()
		console.ShutDownNormal()

	default:
		console.ShowMessage("Command not defined. Check menu for available commands.")
	}
}
