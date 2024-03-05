package main

import "fmt"

func main() {
	s := make([]int, 2, 8)
	fmt.Println(s)

	s = append(s, 42)
	s[0] = 33
	fmt.Println(s)
}
