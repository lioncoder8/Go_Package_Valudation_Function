package main

import (
	"fmt"
	"math/cmplx"
)

was(
	ToBe bool = false
	MaxInt unit64 = 1<<64-1
	z complex128 = cmplx.Sqrt(-5+12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v",ToBe,ToBe)

	fmt.Printf()
}
