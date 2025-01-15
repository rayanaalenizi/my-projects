package main

import (
	"log"
	"os"
)

func broadcast(message string) {
	mu.Lock()
	defer mu.Unlock()
	
	previousMessages = append(previousMessages, message)
	writeToLogFile(message)

	for conn := range clients {
		_, err := conn.Write([]byte(message + "\n"))
		if err != nil {
			log.Printf("Error sending message to client: %v", err)
		}
	}
}

func writeToLogFile(message string) {
	file, err := os.OpenFile("logs/chat_log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Error opening log file: %v", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(message + "\n")
	if err != nil {
		log.Printf("Error writing to log file: %v", err)
	}
}
