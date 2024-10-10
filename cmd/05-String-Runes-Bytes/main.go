package main

import (
	"fmt"
	"strings"
)

// String, Runes and Bytes
//[]rune as alias int32
 func main() {
	 var myString = "resume"
	 var indexed = myString[0]
	 fmt.Printf("%v, %T \n", indexed, indexed)
	 //114, uint8 
	 for i, v := range myString{
		fmt.Println(i, v)
			//0 114
			//1 101
			//2 115
			//3 117
			//4 109
			//5 101
	 }
	
	 var MyRune = 'a' // "string" 'rune'
	fmt.Printf("\n myRune = %v", MyRune)
	var strSlice = []string{"s","u","b","s", "c", "r","i", "b","e"}
	 var catStr = ""
	 for i := range strSlice{
		catStr += strSlice[i]
	 }
	 fmt.Printf("\n %v", catStr)
	 //cannot modified them when created like catStr[0] = "a"
	 // import strings buildstring. string can allocate
	 // much faster
	 var strSlice2 = []string{"s","u","b","s", "c", "r","i", "b","e"}
	var strBuilder strings.Builder
	 for i:= range strSlice2{
		 strBuilder.WriteString(strSlice2[i])
	 }
	  var catStr2 = strBuilder.String()
	fmt.Printf("\n %v", catStr2)

 }