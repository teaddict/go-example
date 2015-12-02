package main

import "fmt"

func Solution(A,B,K int) int {
    var count int = 0;
	if A%K ==0{
		count++
	}
	
	return (B/K) - (A /K) + count	
}

func main() {
	var A int = 6
	var B int = 11
	var K int = 2
	fmt.Println(Solution(A, B, K))
}
