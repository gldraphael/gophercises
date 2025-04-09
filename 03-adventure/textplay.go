package main

import (
	"bufio"
	"fmt"
	"os"
)

func TextModeGamePlay(story AdventureStory) {

	stdin := bufio.NewReader(os.Stdin)

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
				_, err := fmt.Fscan(stdin, &action)
				if err != nil {
					// make sure we skip all input characters until \n
					// to handle cases where a user enters a word instead of a number
					stdin.ReadString('\n')
				}
			}
			currentArc = story.StoryArcs[currentArc.Options[action-1].Arc]
		} else {
			quit = true
		}
	}
}
