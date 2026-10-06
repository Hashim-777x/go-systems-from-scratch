package main 

func main(){
    mux:= http.Router()
	addr:= ":8090"
	http.ListenAndServe(mux,addr)
	we_slice := []int{10, 20, 30} // testing Slice 

	// We use reflect to look at the SliceHeader structure
	header := (*reflect.SliceHeader)(unsafe.Pointer(&we_slice))

	fmt.Printf("Slice Content: %v\n", we_slice)
	fmt.Printf("Pointer to Array: %X\n", header.Data)
	fmt.Printf("Length: %d\n", header.Len)
	fmt.Printf("Capacity: %d\n", header.Cap)
}
