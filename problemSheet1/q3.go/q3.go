


//Taken from  https://play.golang.org/p/Vq-TIHJCHE

package main 

import "fmt"


var text string

func main() {

    fmt.Print("Enter text: ")
    var input string
    fmt.Scanln(&input)

    
     fmt.Print(Reverse(input))


}
   func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}



