package main

import "fmt"

func main() {
	var a = []int{2, 1, 4, 3, 5, 3, 5, 6}
	fmt.Println(a)
	for i, e := range a {
		fmt.Println("i: ",i,"a: ", e)
	}

	for _,e := range a{
		fmt.Println("a: ",e)
	}

	var x uint = 5

	fmt.Println(x)
}

