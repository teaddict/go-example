package main

// you can also use imports, for example:
import "fmt"
import "math"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(X int, Y int, D int) int {
    // write your code in Go 1.4
    var jump int =int(math.Ceil((float64(Y)-float64(X))/float64(D)))
    return jump
}

func main() {
    fmt.Println("result %v\n",Solution(10,85,30))    
}
