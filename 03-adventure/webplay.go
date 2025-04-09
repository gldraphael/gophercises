package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func WebModeGamePlay(story AdventureStory) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, fmt.Sprintf("/%s", story.FirstArc), http.StatusTemporaryRedirect)
	})

	tmpl := template.Must(template.ParseFiles("./template.html"))
	for key, arc := range story.StoryArcs {
		http.HandleFunc(fmt.Sprintf("/%s", key), func(w http.ResponseWriter, r *http.Request) {
      tmpl.Execute(w, arc)
		})
	}


	log.Println("Listening at localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

