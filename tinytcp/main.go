package main

import (
	"bufio"
	"fmt"
	"net"
)

const port = "8080"

func main() {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error starting TCP server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server listening on :" + port)

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Error accepting connection:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Client connected!")

	reader := bufio.NewReader(conn)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading from client:", err)
		return
	}
	fmt.Printf("Received: %s", line)

	_, err = fmt.Fprintln(conn, "I got your message")
	if err != nil {
		fmt.Println("Error writing to client:", err)
		return
	}
}
