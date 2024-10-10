package main

import "fmt"

// Pointers
// what us a pointer
// special type. store mmemort location



func main() {
	 var p *int32 = new(int32) // store init memory in other location but inside pointer int32 location memory 
	  // pointer to int32 location memory still nill 
	 var i int32 // 0 init create store in memory

	 fmt.Printf("The value p points to is: %v", *p) // value 0
	 fmt.Printf("\n The Value if i is : %v", i) // value 0
	 // set and change value memory var use pointer reference *
	 *p = 10
	 fmt.Printf("\nThe value p points to is: %v", *p) // value 10
	 // to use memory other location already init create ,just use reference "&"
	 //refer this location memory
	  p = &i
	  *p =  1
	  fmt.Printf("The value p points to is: %v", *p) // value 1
	  fmt.Printf("\n The Value if i is : %v\n", i) // value 1

	  // other fact. slice has pointer if u copy other slice

	  var slice = []int32{1,2,3}
	 var sliceCopy = slice
	  sliceCopy[2] = 4
	   fmt.Println(slice) // [1,2,4]
	   fmt.Println(sliceCopy) // [1,2,4]

	   var thing1 =  [5]float64{1,2,3,4,5}
	  fmt.Printf("\n The memory of the thing1 array is: %p", &thing1)
	  var result [5]float64 = square(&thing1)
	  fmt.Printf("\nThe result is: %v", result)
	  fmt.Printf("\nThe value of thing1 is: %v", thing1)
}


func square(thing2 *[5]float64) [5]float64 {
	 for i := range thing2{
		 thing2[i] = thing2[i]*thing2[i]
	 }
	 return *thing2
}