package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type RedirectionConfig struct {
	Redirects []struct {
			Slug string
			Url  string
	}
}

func main() {

	configFilePath := flag.String("file", "./config.yaml", "specify a path to the config file")
	flag.Parse()

	var config RedirectionConfig
	var err error
	fileFormat := filepath.Ext(*configFilePath)
	switch fileFormat {
	case ".yaml":
		fallthrough
	case ".yml":
		config, err = getYaml(*configFilePath)
	case ".json":
		config, err = getJson(*configFilePath)
	default:
		log.Fatal("Invalid file format ", fileFormat, ". Only yaml and json files are supported.")
		return
	}
	if err != nil {
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

func getYaml(filePath string) (RedirectionConfig, error) {

	fileText, err := getFileText(filePath)
	if err != nil {
		return RedirectionConfig{}, err
	}

	var config RedirectionConfig
	if err := yaml.Unmarshal([]byte(fileText), &config); err != nil {
        log.Fatal(err)
		return RedirectionConfig{}, err
	}
	return config, nil
}

func getJson(filePath string) (RedirectionConfig, error) {

	fileText, err := getFileText(filePath)
	if err != nil {
		return RedirectionConfig{}, err
	}

	var config RedirectionConfig
	if err := json.Unmarshal([]byte(fileText), &config); err != nil {
        log.Fatal(err)
		return RedirectionConfig{}, err
	}
	return config, nil
}

func getFileText(filePath string) (string, error) {
	fileText, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return string(fileText), nil
}
