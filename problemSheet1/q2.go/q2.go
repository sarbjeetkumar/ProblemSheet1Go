package main

import "fmt"

func main() {

	var number int 
	//var count int 
	//loop through each number one by one 		
	for i := 1; i < 1000; i++ {
		
		for j:= 1 ;j< 1000; j++{

			if(i % j == 1){
				number = i
				//fmt.Println(i)
			}

		}//inner ends here 
	}//for ends here 

	fmt.Println(number)
}




