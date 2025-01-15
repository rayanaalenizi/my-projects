package main

import (
	"log"
	"net"
	"os"
	"sync"
)

var (
	clients          = make(map[net.Conn]string)
	clientColors     = make(map[string]string)
	mu               sync.Mutex
	maxClients       = 10
	previousMessages []string
)

func main() {
	port := "8989"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	defer listener.Close()

	log.Printf("Listening on port :%s", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		mu.Lock()
		if len(clients) >= maxClients {
			mu.Unlock()
			conn.Close()
			continue
		}
		mu.Unlock()

		go handleConnection(conn)
	}
}
