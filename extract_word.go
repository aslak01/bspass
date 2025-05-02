package main

import (
	"bufio"
	"embed"
	"fmt"
	"math/rand"
	"strings"
)

//go:embed words/*
var wordFiles embed.FS

func getRandomWordFromFile(number int) (string, error) {
	filePath := fmt.Sprintf("words/words_%d", number)
	file, err := wordFiles.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("error opening file %s: %w", filePath, err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			lines = append(lines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading file %s: %w", filePath, err)
	}

	if len(lines) == 0 {
		return "", fmt.Errorf("file %s is empty or contains only whitespace", filePath)
	}

	randomIndex := rand.Intn(len(lines))
	return lines[randomIndex], nil
}
