package main

import (
	"fmt"
	"math"
)

const s string = "jimmyio"

func main() {
	fmt.Println(s)

	const n = 50000000000
	const d = 3e20 / n

	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Print(math.Sin(n))

}
