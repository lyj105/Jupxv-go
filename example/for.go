package main

import (
	"fmt"
)

func main() {
	i := 3
	for i < 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 10; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("mxhs")
		break
	}

	for n := 0; n < 7; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}
