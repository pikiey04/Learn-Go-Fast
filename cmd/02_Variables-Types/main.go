package main

import (
	"fmt"
	"unicode/utf8"
)

// Constants variables and Basic Data types

func main() {
	 var intNum int16 = 32767 // bit interger 16 bit
	 intNum = intNum + 1
	 fmt.Println(intNum)

	 var floatNum float64 = 123789.123
	 fmt.Println(floatNum)
	 
	 var floatNum32 float32 = 10.1
	 var intNum32 int32 = 2
	 var result float32 = floatNum32 + float32(intNum32) /// need convert with same type
	  fmt.Println(result)

	  var intNum1 int = 3
	  var intNum2 int = 2
	  fmt.Println(intNum1/intNum2) // 1 no decimal float 
	  fmt.Println(intNum1%intNum2) // remaining 1

	  var myString string =  "hello" + " " + "World" // single line	 
	  fmt.Println(myString) // hello World

	  fmt.Println(len("test")) // 4 length
	  fmt.Println((utf8.RuneCountInString("y")))

	   
	var myRune rune = 'a'
	fmt.Println(myRune) // 97

	var myBoolean bool = false
	 fmt.Println(myBoolean) // false default
	var intNum3 int 
	fmt.Println(intNum3) // 0 default
	myVar := "Text" //  := short hand equal to var. no need sign var
	fmt.Println(myVar) // Text
	// string type default is "" empty string

	// sign multiple variable in single line
	var1 , var2 := 1, 2
	 fmt.Println(var1, var2) // 1 2
	
	// good practice use type safe.. instead use explit := dont know what type
	// add the declare type is good

	 const myConst string = "const value" // const is good choice to sign variable cannot change or modifed
	 fmt.Println(myConst)
	  
}