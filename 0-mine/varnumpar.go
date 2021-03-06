package main

import "fmt"

func main() {
	x := min(1, 3, 2, 0)
	fmt.Printf("The minimum is: %d\n", x)
	arr := []int{7, 9, 3, 5, 1}
	x = min(arr...)
	fmt.Printf("The minimum in the array arr is: %d\n", x)
}

func min(a ...int) int {
	if len(a) == 0 {
		return 0
	}
	min_v := a[0]
	for _, v := range a {
		if v < min_v {
			min_v = v
		}
	}
	return min_v
}

