package main

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func readHighScore() int {
	path := "data/.snake_highscore"
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("High score file not found at %s, starting with 0", path)
			return 0
		}
		log.Printf("Error reading high score file at %s: %v", path, err)
		return 0
	}
	scoreStr := strings.TrimSpace(string(data))
	score, err := strconv.Atoi(scoreStr)
	if err != nil {
		log.Printf("Error parsing high score from %s: %v", path, err)
		return 0
	}
	return score
}

func writeHighScore(score int) {
	dir := "data"
	path := filepath.Join(dir, ".snake_highscore")
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		log.Printf("Error creating directory %s: %v", dir, err)
		return
	}
	err = os.WriteFile(path, []byte(strconv.Itoa(score)), 0644)
	if err != nil {
		log.Printf("Error writing high score to %s: %v", path, err)
	}
}
