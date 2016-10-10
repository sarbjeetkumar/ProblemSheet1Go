 package main

//this progrm prints the values of collatz cojection


import "fmt"

func collatz(n unit){

	for : n! = 1 {

	if n % 2 == 0{

	n = n/ 2
	}else{
	n = (3*n) +1
	}

}

	fmt.println(n)
}


func main(){
	collatz(20)
	collatz(400)


}
