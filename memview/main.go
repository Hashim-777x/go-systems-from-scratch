package main

import (
	"fmt"
	"unsafe"
)

type BadLayout struct {
	A bool   // 1 byte and padding +7 bytes to make it -> 8 bytes for formating
	B int64 // 8 bytes it takes 
	C bool // 1 byte and + 7 bytes of trailing padding ! -> 8 bytes 
}       // -------------------------------TOTAL SIZE = 24 BYTES --------------------------

type GoodLayout struct {
	B int64  // 8 bytes ( may takes from 0 - 7 bytes )
	A bool  // 1 bytes 
	C bool  // 1 bytes  // 1 + 1 bytes -> 2 bytes and 6 trailing bytes -> 8 bytes 
}

func main() {
	bad := BadLayout{}
	good := GoodLayout{}

	fmt.Println("BadLayout size: ", unsafe.Sizeof(bad))
	fmt.Println("  A offset:", unsafe.Offsetof(bad.A), "  (size", unsafe.Sizeof(bad.A), ")")
	fmt.Println("  B offset:", unsafe.Offsetof(bad.B), "  (size", unsafe.Sizeof(bad.B), ")")
	fmt.Println("  C offset:", unsafe.Offsetof(bad.C), "  (size", unsafe.Sizeof(bad.C), ")")

	fmt.Println()

	fmt.Println("GoodLayout size:", unsafe.Sizeof(good))
	fmt.Println("  B offset:", unsafe.Offsetof(good.B), "  (size", unsafe.Sizeof(good.B), ")")
	fmt.Println("  A offset:", unsafe.Offsetof(good.A), "  (size", unsafe.Sizeof(good.A), ")")
	fmt.Println("  C offset:", unsafe.Offsetof(good.C), "  (size", unsafe.Sizeof(good.C), ")")
}

