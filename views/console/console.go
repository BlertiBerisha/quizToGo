package console

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"quizToGo/models"
	"sort"
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

// ShowAvailableTopics displays the available topics
func ShowAvailableTopics() {
	fmt.Println("Available topics: Physics, History, Mathematics")
	fmt.Println("Type the topic you want to play with or type 'Random'")
}

// ShowAvailableDifficulties displays the available difficulty levels
func ShowAvailableDifficulties() {
	fmt.Println("Available difficulties: Easy, Medium, Hard")
}

// ShowHowManyQuestions asks the user how many questions they would like to answer
func ShowHowManyQuestions() {
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

// ShowHighScores displays all high scores sorted by success rate
func ShowHighScores(scores []models.Score) {
	// Sort scores by success rate (descending)
	sort.Slice(scores, func(i, j int) bool {
		// Calculate success rate for both scores
		rateI := float64(scores[i].Score) / float64(scores[i].TotalQuestions)
		rateJ := float64(scores[j].Score) / float64(scores[j].TotalQuestions)
		return rateI > rateJ // Sort in descending order
	})

	// Display sorted high scores
	fmt.Println("\nHIGH SCORES:")
	fmt.Println("-------------------")
	for i, score := range scores {
		rate := float64(score.Score) / float64(score.TotalQuestions) * 100
		fmt.Printf("%d. %s: %d/%d (%.1f%%)\n",
			i+1,
			score.PlayerName,
			score.Score,
			score.TotalQuestions,
			rate,
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
