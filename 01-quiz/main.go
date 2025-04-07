package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
    fmt.Println("hello world")

	const CSV_FILE_PATH = "problems.csv"
	csv_file, err := os.Open(CSV_FILE_PATH)
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(csv_file)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
		return
	}

	total_questions := len(records)
	score := 0
	for _, record := range records {
		question := record[0]
		answer := record[1]
		var input string

		fmt.Printf("%s = ", question)
		fmt.Scan(&input)

		if input == answer {
			score++
		}
	}

	fmt.Println()
	fmt.Println("Your score is", score, "/",total_questions)
	fmt.Println("You got",(float32(score) / float32(total_questions))*100, "% questions right.")
}
