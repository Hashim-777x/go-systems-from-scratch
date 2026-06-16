package main 

import (
	"fmt"
	"net"
)
 
func main() {
	fmt.Println("Hello, World!")
    const PORT = "8080"
	
	listener, err := net.Listen("tcp", "localhost:"+PORT)
	if err != nil {
		fmt.Println("Error starting TCP server:", err)
		return
	}
	defer listener.Close()

}
