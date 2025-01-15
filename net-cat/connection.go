package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)


func handleConnection(conn net.Conn) {
	defer conn.Close() 


	conn.Write([]byte(logo() + "\n\n[ENTER YOUR NAME]: "))
	name := "" 

	for {
		
		reader := bufio.NewReader(conn)
		input, err := reader.ReadString('\n')
		if err != nil {
		
			break
		}
		
		input = strings.TrimSpace(input)

		if name == "" {
			
			if input == "" {
				conn.Write([]byte("Name cannot be empty, please enter your name: "))
				continue 
			}
			name = input 

			mu.Lock()
			clients[conn] = name
			clientColors[name] = assignColor(name)
			mu.Unlock()

	
			clearLastLine(conn)

			broadcast(fmt.Sprintf("%s%s has joined the chat...\033[0m", clientColors[name], name))

			
			sendPreviousMessages(conn, name)

			
			// displayPrompt(conn, name)
			continue
		}

		if input != "" {
			clearLastLine(conn) 
			broadcast(fmt.Sprintf("[%s][%s%s\033[0m]: %s", time.Now().Format("2006-01-02 15:04:05"), clientColors[name], name, input))
		}

		// displayPrompt(conn, name)
	}

	
	mu.Lock()
	delete(clients, conn)
	mu.Unlock()

	broadcast(fmt.Sprintf("%s%s has left the chat...\033[0m", clientColors[name], name))
}

func clearLastLine(conn net.Conn) {
	_, _ = conn.Write([]byte("\033[F\033[K")) 
}

// func displayPrompt(conn net.Conn, name string) {
// 	_, _ = conn.Write([]byte(fmt.Sprintf("[%s][%s]: ", time.Now().Format("2006-01-02 15:04:05"), name)))
// }
