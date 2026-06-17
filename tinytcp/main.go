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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	}
	http.ListenAndServe("localhost:"+PORT, nil)

	fmt.Println("TCP server listening on port", PORT)

}
