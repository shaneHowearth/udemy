package main

import "fmt"

func main() {
	sl1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range sl1 {
		if v%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}
}
