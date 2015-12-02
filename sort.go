package main

import(
	"fmt"
	"sort"
)

func main() {
	x := []int{4,2,1,3,5,6}
	sort.Ints(x)
	fmt.Println(x)

	y := []int{3,2,1,5,2,3,5,6}
	sort.Ints(y)
	fmt.Println()y
}