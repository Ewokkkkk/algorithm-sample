package main

import "fmt"

func main() {
	s := []int{2, 4, 2, 1, 0, 3, 6, 5, 10, 9, 8, 12, 11}
	InsertionSort(s)
	fmt.Println(s)
}

func InsertionSort(s []int) []int {
	for i := 1; i < len(s); i++ {
		for j := 0; j < i; j++ {
			if s[i-j-1] > s[i-j] {
				s[i-j-1], s[i-j] = s[i-j], s[i-j-1]
			} else {
				break
			}
		}
	}
	return s
}
