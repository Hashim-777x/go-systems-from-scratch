package main

import (
	"fmt"
	"unsafe"
)

type BadLayout struct {
	A bool
	B int64
	C string
}
type GoodLayout struct {
	A bool
	B int64
	C string
}

func main() {
	Bad:= BadLayout{}
	fmt.Println( "bad lay out size : ", unsafe.Sizeof(Bad) )
	
}