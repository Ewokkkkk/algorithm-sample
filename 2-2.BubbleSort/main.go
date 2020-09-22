package main

import "fmt"

func main() {
	s := []int{2, 4, 2, 1, 0, 3, 6, 5, 10, 9, 8, 12, 11}
	BubbleSort(s)
	fmt.Println(s)
}

func BubbleSort(s []int) []int {
	for i := 0; i < len(s)-1; i++ {
		for j := len(s) - 1; j > i; j-- {
			if s[j-1] > s[j] {
				s[j-1], s[j] = s[j], s[j-1]
			}
		}
	}
	return s
}
