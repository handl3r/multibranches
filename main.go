package main

import "fmt"

func main() {
	fmt.Println("Hello, this is multibranches app")
	fmt.Printf("Result of call Min(2, 3) is: %d\n", Min(2, 3))
	fmt.Println("End")
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
