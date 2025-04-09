package main

import (
	"encoding/json"
	"os"
)

type StoryArcOption struct {
	Text string
	Arc  string
}

type StoryArc struct {
	Title   string
	Story   []string
	Options []StoryArcOption
}

type AdventureStory struct {
	FirstArc  string
	StoryArcs map[string]StoryArc
}

func GetStory(filePath string) (AdventureStory, error) {

	jsonText, err := os.ReadFile(filePath)
	if err != nil {
		return AdventureStory{}, err
	}

	story := AdventureStory{}
	json.Unmarshal([]byte(jsonText), &story)
	return story, nil
}

