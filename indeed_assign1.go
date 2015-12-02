// INDEED PRIME CODE CHALLENGE
// NOVEMBER 21-22 , 2015
// %100

package main

import (
    "fmt"
    "strings"
    "regexp"
    "unicode"
)

func Solution(S string) int {
	if len(S)>200 || len(S)<1{
		return -1
	}
    s := strings.Split(S," ")
    var i int = 0
    var n int = len(s)
    var countdigit int = 0
    var countletter int = 0
    found := []int{}
    for i<n{
    	countdigit = 0
    	countletter = 0
    	if regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(s[i]) {
    		fmt.Println(s[i],"ok")
    		for _, value := range s[i] {
 		    	switch{
    				case unicode.IsDigit(value):
        				countdigit++
        			case unicode.IsLetter(value):		
        				countletter++
    			}
			}		
			if(countdigit%2!=0 && countletter%2==0){
				fmt.Println(countdigit,countletter)
				found = append(found,len(s[i]))
			}
		}
		i++
	}
	if len(found) != 0{
		var max int = found[0]
		var i int = 0
		var n int = len(found)

		for i<n{
			if max<found[i]{
				max = found[i]
			}
			i++
		}
		return max
	}else{
    	return -1
    }
}

func main() {
    var S string = "test 5 a0A pass007 ?xy1 +5362- 2222 45643 a"
    fmt.Println(Solution(S))
}