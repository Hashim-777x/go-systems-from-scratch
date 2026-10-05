package main 

func main(){
    mux:= http.Router()
	addr:= ":8090"
	http.ListenAndServe(mux,addr)
}
