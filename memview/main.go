package main

import (
	"fmt"
	"unsafe"
)

type BadLayout struct {
	A bool
	B int64
	C bool
}

type GoodLayout struct {
	B int64
	A bool
	C bool
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

