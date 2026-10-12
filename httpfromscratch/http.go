package main 
"net/http"

func main(){
    mux:=http.NewServeMux()
	address:= ":8949"
	http.listenAndServer(mux,address)
	
}