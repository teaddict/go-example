// INDEED PRIME CODE CHALLENGE
// NOVEMBER 21-22 , 2015
// %75

package main
import "fmt"

func Solution(A []int) int {
    // write your code in Go 1.4
    var n int = len(A)
    var i int = 0
    var lastmax int = A[0]
    var firstmax int = 0
    var max int = 0
    for i<n-1{
    	if i==0{
    		if A[0]<A[1]{
    			i++
    			continue
    		}
    	}
    	if i==n-2{
    		if(A[n-1]<A[n-2]){
    			i++
    			continue
    		}
    	}
    	fmt.Println(firstmax,lastmax,A[i])
    	if(lastmax<A[i]){
    		firstmax = lastmax
    		lastmax = A[i]
    		fmt.Println(firstmax,lastmax)
    		if(A[i]>A[i-1]){
    			if(max<A[i]-A[i-1]){
    				fmt.Println(lastmax,firstmax,A[i-1])
    				max = firstmax - A[i-1]
    				fmt.Println(max)
    			} 
    		}
    	}
    	i++
    }
    return max
}

func main() {
	var A = []int{1,3,2,1,2,1,5,3,3,4,2}
	fmt.Println(Solution(A))
}	