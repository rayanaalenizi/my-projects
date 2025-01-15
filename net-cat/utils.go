package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func logo() string {
	purple := "\033[35m"
	pink := "\033[95m"
	reset := "\033[0m"

	file, err := os.Open("assets/logo.txt")
	if err != nil {
		log.Fatalf("Error reading logo file: %v", err)
	}
	defer file.Close()

	var coloredLogo string
	scanner := bufio.NewScanner(file)
	isPink := true

	for scanner.Scan() {
		line := scanner.Text()
		if isPink {
			coloredLogo += fmt.Sprintf("%s%s%s\n", pink, line, reset)
		} else {
			coloredLogo += fmt.Sprintf("%s%s%s\n", purple, line, reset)
		}
		isPink = !isPink
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading logo file: %v", err)
	}

	return coloredLogo
}
