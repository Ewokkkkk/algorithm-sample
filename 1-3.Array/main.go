package main

import "fmt"

func main() {
	arr := []string{"Blue", "Yellow", "Red"}
	fmt.Println(arr)

	arr = apnd(arr, "Green", 1) // [1]にGreenを追加
	fmt.Println(arr)

	arr = dlt(arr, "Green") // Greenを検索し、削除
	fmt.Println(arr)

}
func apnd(l []string, t string, p int) []string {
	l = append(l, "")
	for i := len(l) - 1; i > p; i-- {
		l[i] = l[i-1]
	}
	l[p] = t
	return l
}
func dlt(l []string, t string) []string {
	for j, i := range l {
		if i == t {
			l[j] = ""
			for k := j; k < len(l); k++ {
				if k != len(l)-1 {
					l[k] = l[k+1]
				} else {
					l = l[:k]
				}
			}
		}
	}
	return l
}
