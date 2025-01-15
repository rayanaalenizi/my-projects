package main

import (
	"fmt"
	"net"
	"strings"
)

func sendPreviousMessages(conn net.Conn, name string) {
	mu.Lock()
	defer mu.Unlock()

	for _, msg := range previousMessages {
		if strings.Contains(msg, fmt.Sprintf("%s has joined the chat...", name)) {
			continue
		}
		conn.Write([]byte(msg + "\n"))
	}
}
