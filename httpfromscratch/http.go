package main

import "net/http"
import "fmt"
func main() {
	mux := http.NewServeMux()
	address := ":8090"
	http.ListenAndServe(address, mux)
}
