package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gorilla/websocket"
)

func main() {
	var connections int
	flag.IntVar(&connections, "n", 1, "Specifies the number of open connections")
	flag.Parse()

	// To be productiion ready
	domain := os.Getenv("DOMAIN")
	if domain == "" {
		domain = "localhost:8080"
	}
	url := fmt.Sprintf("ws://%s/goapp/ws", domain)

	wsConnections := make([]*websocket.Conn, connections)
	var wg sync.WaitGroup
	stopCh := make(chan struct{})

	for i := 0; i < connections; i++ {
		wg.Add(1)

		go func(index int) {
			defer wg.Done()

			ws, _, err := websocket.DefaultDialer.Dial(url, nil)
			if err != nil {
				log.Fatalf("Error connecting to the websocket (connection %d): %v", index+1, err)
			}
			defer ws.Close()

			// Store the connection for closing it later
			wsConnections[index] = ws
			fmt.Printf("Connected to WebSocket (connection %d)\n", index+1)

			for {
				select {
				case <-stopCh:
					return
				default:
					_, msg, err := ws.ReadMessage()
					if err != nil {
						log.Printf("Error reading from websocket (connection %d): %v", index+1, err)
						return
					}
					fmt.Printf("RESPONSE from connection %d: %s\n", index+1, msg)
				}
			}
		}(i)
	}

	// Set up signal capturing to handle Ctrl-C
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	fmt.Println("\nReceived interrupt signal, closing connections...")

	close(stopCh)
	for _, ws := range wsConnections {
		if ws != nil {
			ws.Close()
		}
	}

	wg.Wait()
	fmt.Println("All connections closed.")
}
