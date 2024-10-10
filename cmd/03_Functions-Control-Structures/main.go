package main

import (
	"errors"
	"fmt"
)

// function and Control Structures

func main() {
	var printme string = "Hello world"
	printMe(printme) // Hello world
}

// add other function
func printMe(printValue string) {
	 fmt.Println(printValue) 
	var  num , deno int = 11,2
	var result, remainder, err  =  intDivision(num,deno)

	// use switch
	switch {
	case err  != nil: 
	 	fmt.Println(err.Error())
	case remainder==0:
		fmt.Printf("The result of the integer division is %v", result)
	default:
		fmt.Printf("the result of the integer division is %v with remainder %v ", result,remainder)
	}
	switch remainder{
	case 0:
		fmt.Printf("The division was exact")
	case 1,2:
		fmt.Printf("The division was close")
	default:
		fmt.Printf("The division was not close")
	}
	// if err != nil{
	// 	 fmt.Printf(err.Error())
	// }else if remainder == 0{
	// 	 fmt.Printf("The result of the integer division is %v", result)
	// }else {

	// 	fmt.Printf("the result of the integer division is %v with remainder %v", result,remainder)
	// }
	  //printf("string %v", variablevalue)
	   
}
// func Namefunction(params type) type of return {excute stuff function} 
func intDivision(numerator int, denominator int) (int,int,error) {
	// to prevent error or handle error
	var err error
	 if denominator==0 {
		 err = errors.New("cannot divide by zero")
		 return 0,0, err
	 }
	 var result int = numerator/denominator
	 var remainder int = numerator%denominator
	 return result, remainder, err
}