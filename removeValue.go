package main

import "fmt"

func remove(x []int, y int) (a []int) {
	for i := 0; i < len(x); {
		if y == x[i] {
			x[i] = x[i+1]
			a = x[:len(x)-1]
			return a
		}
		i++
	}

	return []int{}
}

func main() {
	s := []int{1, 2, 3, 4}
	fmt.Println(s)
	s = remove(s, 3)
	fmt.Println(s)
}
