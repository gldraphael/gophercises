package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const STORY_FILE_PATH = "./gopher.json"

func main() {
	stdin := bufio.NewReader(os.Stdin)
	story, err := GetStory(STORY_FILE_PATH)
	if err != nil {
		log.Fatal("Could not read", STORY_FILE_PATH)
	}
	
	currentArc := story.StoryArcs[story.FirstArc]
	quit := false
	for quit != true {

		fmt.Println()
		fmt.Printf("⭐ %s\n", currentArc.Title)
		fmt.Println("------------------------------------")
		fmt.Println()
		for _, paragraph := range currentArc.Story {
		  fmt.Printf("  %s\n", paragraph)
		  fmt.Println()
		}
		if len(currentArc.Options) > 0 {
			action := 0
			for action < 1 || action > len(currentArc.Options) {
				fmt.Printf(" ➡️ Pick an action:\n")
				for i, o := range currentArc.Options {
					fmt.Printf(" %d. %s\n", i+1, o.Text)
				}
				fmt.Print("\nYour action: ")
				_, err = fmt.Fscan(stdin, &action)
				if err != nil {
					stdin.ReadString('\n') // skip all input characters until \n
				}
			}
			currentArc = story.StoryArcs[currentArc.Options[action-1].Arc]
		} else {
			quit = true 
		}
	}
}

