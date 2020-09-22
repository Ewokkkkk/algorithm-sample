package main

import "fmt"

func main() {
	s := []int{2, 4, 2, 1, 0, 3, 6, 5, 10, 9, 8, 12, 11}
	SelectionSort(s)
	fmt.Println(s)
}
func Min(l []int) (index, value int) {
	value = l[0]
	index = 0
	for i, v := range l {
		if value > v {
			value = v
			index = i
		}
	}
	return
}

func SelectionSort(s []int) []int {
	for i := 0; i < len(s)-1; i++ {
		index, _ := Min(s[i:])
		s[i], s[i+index] = s[i+index], s[i]
	}
	return s
}
