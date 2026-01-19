package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("CI/CD Demo: 1 + 2 =", Add(1, 2))
}
