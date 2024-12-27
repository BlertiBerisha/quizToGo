package console

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"quizToGo/models"
	"strings"
)

// ExitStatusCodeNoError exit status code no error
const ExitStatusCodeNoError int = 0

// ShowMenu shows the menu to the console
func ShowMenu() {
	fmt.Println(`
	###########################################
	#********** WELCOME TO QUIZ TO GO ********
	#********* CHOOSE YOUR OPTION BELOW ******
	# 1. START NEW QUIZ
	# 2. VIEW HIGH SCORES
	#
	# c. CLEAR VIEW AND SHOW MENU
	# q. QUIT QUIZ GAME
	`)
}

// ShowQuizSetup shows the quiz setup information
func ShowQuizSetup() {
	fmt.Println("How many questions would you like to answer?")
}

// ShowQuestion displays a question and its options
func ShowQuestion(question models.Question) {
	fmt.Printf("\nQuestion: %s\n\n", question.Text)
	for i, option := range question.Options {
		fmt.Printf("%d. %s\n", i+1, option)
	}
	fmt.Print("\nYour answer (1-", len(question.Options), "): ")
}

// ShowCorrectAnswer displays correct answer message
func ShowCorrectAnswer() {
	fmt.Println("\nCorrect! Well done!")
}

// ShowWrongAnswer displays wrong answer message
func ShowWrongAnswer(correctAnswer string) {
	fmt.Printf("\nSorry, that's incorrect. The correct answer was: %s\n", correctAnswer)
}

// ShowFinalScore displays the final score
func ShowFinalScore(score, total int) {
	fmt.Printf("\nYour final score: %d out of %d\n", score, total)
	percentage := float64(score) / float64(total) * 100
	fmt.Printf("Percentage: %.1f%%\n", percentage)
}

// ShowHighScores displays all high scores
func ShowHighScores(scores []models.Score) {
	fmt.Println("\nHIGH SCORES:")
	fmt.Println("-------------------")
	for i, score := range scores {
		fmt.Printf("%d. %s: %d/%d (%.1f%%)\n",
			i+1,
			score.PlayerName,
			score.Score,
			score.TotalQuestions,
			float64(score.Score)/float64(score.TotalQuestions)*100,
		)
	}
}

// ShowContinue shows the continuation prompt
func ShowContinue() {
	fmt.Println("\nPress any key to continue...")
}

// ShowGoodbye shows the goodbye message
func ShowGoodbye() {
	fmt.Println("Thanks for playing! Goodbye!")
}

// ShowMessage shows a message to the console
func ShowMessage(message string) {
	fmt.Println(message)
}

// ShowError shows an error message
func ShowError(err error) {
	if err != nil {
		log.Println("Error:", err.Error())
	}
}

// Clear clears the console view
func Clear() {
	c := exec.Command("cmd", "/c", "cls")
	c.Stdout = os.Stdout
	_ = c.Run()
}

// AskForInput reads user input from console
func AskForInput() string {
	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('\n')
	return strings.TrimSpace(response)
}

// ShutDownNormal terminates the application
func ShutDownNormal() {
	os.Exit(ExitStatusCodeNoError)
}
