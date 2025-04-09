package main

import (
	"log"

	flag "github.com/spf13/pflag"
)

const DEFAULT_STORY_FILE_PATH = "./gopher.json"

func main() {

	storyFilePath := flag.StringP("story", "s", DEFAULT_STORY_FILE_PATH, "A path to the JSON story file.")
	gameplayMode  := flag.StringP("mode",  "m", "text", "The gameplay mode. Supported values: text, web")
	flag.Parse()

	if *gameplayMode != "text" && *gameplayMode != "web" {
		log.Fatalf("Invalid value '%s' for --mode. Supported values: text, web\n", *gameplayMode)
	}

	story, err := GetStory(*storyFilePath)
	if err != nil {
		log.Fatalln("Could not read", DEFAULT_STORY_FILE_PATH)
	}

	switch *gameplayMode {
		case "text": TextModeGamePlay(story)
		case "web":  WebModeGamePlay(story)
	}
	
}

