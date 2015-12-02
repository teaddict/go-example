package main

// you can also use imports, for example:
import "fmt"
import "sort"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")
func Solution(A []int) int {
    // write your code in Go 1.4
    var max int = len(A)
    var i int = 0
    sort.Ints(A)
    if max == 0 {
        return 1
    }
    for i<max {
        if i+1 != A[i] {
            return i+1
        }else{
            i = i +1
        }
    }
    return 0
}

func main() {
    var A = []int {2,3,1,5}
    fmt.Println("result: ",Solution(A))
    var B = []int {2,4,1,5}
    fmt.Println("result: ",Solution(B))
    var C = []int {}
    fmt.Println("result: ",Solution(C))
}
