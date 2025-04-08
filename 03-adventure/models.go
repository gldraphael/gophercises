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

// type Story struct {
// 	Intro     StoryArc `json:"intro"` 
// 	NewYork   StoryArc `json:"new-york"`
//   Debate    StoryArc `json:"debate"`
//   SeanKelly StoryArc `json:"sean-kelly"`
//   MarkBates StoryArc `json:"mark-bates"`
//   Denver    StoryArc `json:"denver"`
// 	Home      StoryArc `json:"home"`
// }

type AdventureStory struct {
	FirstStoryArc string
	StoryArcs     map[string]StoryArc
}

func GetStory(filePath string) (AdventureStory, error) {

	jsonText, err := os.ReadFile(filePath)
	if err != nil {
		return AdventureStory{}, err
	}

	story := AdventureStory{ FirstStoryArc: "intro" }
	json.Unmarshal([]byte(jsonText), &story.StoryArcs)
	return story, nil
}

