package main

import (
	"fmt"
	"log"
)

const STORY_FILE_PATH = "./gopher.json"

func main() {
	story, err := GetStory(STORY_FILE_PATH)
	if err != nil {
		log.Fatal("Could not read", STORY_FILE_PATH)
	}
	
	currentArc := story.StoryArcs[story.FirstStoryArc]
	quit := false
	for quit != true {

		fmt.Println()
		fmt.Println(currentArc.Title)
		fmt.Println("------------------------------------")
		fmt.Println()
		for _, paragraph := range currentArc.Story {
		  fmt.Printf("  %s\n", paragraph)
		  fmt.Println()
		}
		if len(currentArc.Options) > 0 {
			fmt.Println(" ➡️ Pick an action:")
			for i, o := range currentArc.Options {
				fmt.Printf(" %d. %s\n", i+1, o.Text)
			}
			fmt.Print("\nYour action: ")
			var action int
			fmt.Scan(&action)
			// TODO: add validation
			currentArc = story.StoryArcs[currentArc.Options[action-1].Arc]
		} else {
			quit = true 
		}
	}
}

