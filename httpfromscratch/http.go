package main

import (
	"net/http"
	"fmt"
)
func main() {
	fmt.Println("HTTP from scratch")
	mux := http.NewServeMux()
	address := ":8090"
	http.ListenAndServe(address, mux)
}
