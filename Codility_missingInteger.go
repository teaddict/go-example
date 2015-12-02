package solution
//1....N dizisinde hangisi eksik onu buluyor
// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
    // write your code in Go 1.4
    var sum int =0
    var i int =0
    var Nsum int = 0
    var N int = len(A)
    for i<N{
        sum = sum +A[i]
        i++
    }
    if (N%2==0){
        Nsum = (N+1)*((N+2)/2)
        return int (Nsum-sum)
    }else{
        Nsum=((N+1)/2)*(N+2)
        return int(Nsum-sum)
    }
}

