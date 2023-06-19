package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type Question struct {
	Text   string
	Answer string
}

func main() {
	// Define flags
	filePath := flag.String("csv", "problem.csv", "Path to the CSV file")
	timeLimit := flag.Int("time", 15, "Time limit (in seconds) for the quiz")
	flag.Parse()

	// Check if the file flag is provided
	if *filePath == "" {
		fmt.Println("Please provide the path to a CSV file using the -file flag.")
		return
	}

	// Define timer
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	// Open the CSV file
	file, err := os.Open(*filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read all the CSV records
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	// Press Enter to start
	fmt.Println("Press enter to start the quiz")
	fmt.Scanln()

	// Process each record and ask questions
	correctCount := 0
QuestionLoop:
	for i, record := range records {
		// Create a new Question
		question := Question{
			Text:   record[0],
			Answer: record[1],
		}

		fmt.Println("Question", i+1, ":", question.Text)

		answerCh := make(chan string)

		go func() {
			var userAnswer string
			fmt.Print("Your Answer: ")
			fmt.Scan(&userAnswer)
			answerCh <- userAnswer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break QuestionLoop

		case userAnswer := <-answerCh:
			userAnswer = strings.TrimSpace(strings.ToLower(userAnswer))
			correctAnswer := strings.TrimSpace(strings.ToLower(question.Answer))

			if userAnswer == correctAnswer {
				correctCount++
			}
		}
		fmt.Println()

	}

	fmt.Printf("You answered %d out of %d questions correctly.\n", correctCount, len(records))
}
