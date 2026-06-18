package main

import (
	"bufio"
	"fmt"
	"net"
)

const port = "8080"

func main() {
	conn, err := net.Dial("tcp", "localhost:"+port)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	_, err = fmt.Fprintln(conn, "hello from client")
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}

	reader := bufio.NewReader(conn)
	reply, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading server response:", err)
		return
	}

	fmt.Printf("Server says: %s", reply)
}
