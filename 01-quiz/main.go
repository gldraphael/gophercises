package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	time_limit := flag.Int("timer", 2, "Default time limit for each question.")

	const CSV_FILE_PATH = "problems.csv"
	records, err := load_questions(CSV_FILE_PATH)
	if err != nil {
		log.Fatalln("Questions could not be loaded.", err)
	}
	
	print_welcome_prompt()
	

	total_questions := len(records)
	score := 0
	for _, record := range records {
		question := record[0]
		answer := record[1]

		timer := time.NewTimer(time.Duration(*time_limit) * time.Second)

		fmt.Print(timer.C, "➡️ ", question, " = ")
		inputCh := make(chan string)
		go func() {
			var input string
			fmt.Scan(&input)
			inputCh <- input
		}();

		select {
		case <-timer.C:
			println("⏲️", timer.C)
		case input := <- inputCh:
			timer.Stop()
			if input == answer {
				score++
			}
		}
	}

	print_scores(score, total_questions)
}

func print_welcome_prompt() {

	fmt.Println("Welcome to the goquiz app!")
	fmt.Println()

	fmt.Println("How to play")
	fmt.Println("---------------------------")
	fmt.Println("You will be asked a few questions.\nAnswer them as fast as you can.\nPress ENTER to proceed to the next question.")
	fmt.Println()
	fmt.Println("Press ENTER to begin!")
	fmt.Scanln()
}

func print_scores(score int, total_questions int) {
	fmt.Println()
	fmt.Println("Your score is", score, "/",total_questions)
	fmt.Println("You got",(float32(score) / float32(total_questions))*100, "% questions right.")
}

func load_questions(csv_path string) ([][]string, error) {
	csv_file, err := os.Open(csv_path)
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(csv_file)
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}
