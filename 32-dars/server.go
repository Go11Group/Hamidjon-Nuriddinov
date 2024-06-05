package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

var (
	clients    = make(map[net.Conn]bool)
	broadcast  = make(chan string)
	register   = make(chan net.Conn)
	unregister = make(chan net.Conn)
	mu         sync.Mutex
)

func handleConnection(conn net.Conn) {
	defer func() {
		unregister <- conn
		conn.Close()
	}()

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		clientAddr := conn.RemoteAddr().String()
		formattedMessage := fmt.Sprintf("%s: %s", clientAddr, message)
		broadcast <- formattedMessage
	}
}

func manageClients() {
	for {
		select {
		case conn := <-register:
			mu.Lock()
			clients[conn] = true
			mu.Unlock()
			fmt.Println("New client connected:", conn.RemoteAddr())
		case conn := <-unregister:
			mu.Lock()
			if _, ok := clients[conn]; ok {
				delete(clients, conn)
				fmt.Println("Client disconnected:", conn.RemoteAddr())
			}
			mu.Unlock()
		case message := <-broadcast:
			fmt.Print("Message received: ", message) // Ma'lumotni serverda ko'rsatish
			mu.Lock()
			for client := range clients {
				if _, err := client.Write([]byte(message)); err != nil {
					client.Close()
					delete(clients, client)
				}
			}
			mu.Unlock()
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting TCP server:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening on port 8080")

	go manageClients()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		register <- conn
		go handleConnection(conn)
	}
}
