package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	fmt.Println("s ==", s)
	fmt.Println("s[1:4] ==", s[1:4])
	fmt.Println("s[1:] ==", s[1:])
	fmt.Println("s[:4] ==", s[:4])

}