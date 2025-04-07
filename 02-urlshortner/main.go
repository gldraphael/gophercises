package main

import (
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

type RedirectionConfig struct {
	Redirects []struct {
			Slug string
			Url  string
	}
}

func main() {

	fileText, err := os.ReadFile("./config.yaml")
	if err != nil {
		log.Fatal(err)
		return
	}

	var config RedirectionConfig
	if err := yaml.Unmarshal([]byte(fileText), &config); err != nil {
        log.Fatal(err)
		return
	}

	for _, rule := range config.Redirects {
		http.HandleFunc(rule.Slug, func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, rule.Url, http.StatusTemporaryRedirect)
		})
	}

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
