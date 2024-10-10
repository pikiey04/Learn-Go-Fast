package main

import (
	"fmt"
	"time"
)

// Arrays, Slices, Maps
// Looping Control Structures
// []Arrays = fixed Length - Same Type - indexable -Contiguous in Memory

func main() {
	var  intArr [3]int32
	intArr[1] = 123 // only can  use same type
	intArr[2] = 321
	fmt.Println(intArr[0]) // indexable
	fmt.Println(intArr[1])
	fmt.Println(intArr[2])
	fmt.Println(intArr[0:3]) // [index, end of length]
	fmt.Println(intArr[2:3])
	fmt.Println(&intArr[2]) // at location address memory 0x14000112028 for index 2 length 3 
 // in short hand u can state array like this
  intArr1 := [3]int32{1,2,3} // or
  intArr2 := [...]int32{1,2,3,4,5,6} // like this
  fmt.Println(intArr1)
  fmt.Println(intArr2)


  // Slice wrap arrays to give a more general, powerful, and convenient
  // general, powerful, and convenient in terfae to sequences of data

   var intSlice []int32 = []int32{4,5,6}
	fmt.Println(intSlice) // [4 5 6]	
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7) //add value end of array with new capacity
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))
	fmt.Println(intSlice) // [4 5 6 7, *, *]  store duplacte cap from last
	
	var intSlice2 []int32 = []int32{8,9}
	 intSlice = append(intSlice, intSlice2...)
	 fmt.Println(intSlice) // [4 5 6 7 8 9]
	 fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))

	 // other way to make slice 
	
	var intSlice3 []int32 = make([]int32,3,8)
	fmt.Println(intSlice3) // [0 0 0 * * * * *]
	fmt.Printf("The length is %v with capacity %v", len(intSlice3), cap(intSlice3))

	  // map[string]int32
	  // key value pair {"key": "pair"}

	   var myMap map[string]uint8 = make(map[string] uint8)
	   fmt.Println(myMap)
	   var myMap2 = map[string]uint8{"Adam":23, "Sarah": 45, "John": 55}
	   fmt.Println(myMap2["Adam"]) // 23
	   fmt.Println(myMap2["Jason"]) // 0 even key doesnt exist
	var age, ok = myMap2["Jason"]  // second argument will be boolean true or false . if in map doesnt have key on age then will be return false in ok 
	 if ok{
		fmt.Printf("The age is %v", age)

	 }else{ 
		 fmt.Println("Invalid Name")
	 }
     delete(myMap2, "Adam") // (map, key) to delete from map
	 fmt.Println(myMap2) // map[Sarah:45 John:55]

	// for loop

	for name, age := range myMap2{
		fmt.Printf("Name: %v, Age:%v \n", name,age)
	}
	//Name: Sarah, Age:45 
	//Name: John, Age:55 
 // for init; condition; excutelooppost{}
  for i:=0; i<10; i++ {
	 fmt.Println(i)
  } 

   var n int = 1000000
var testSlice = []int{}
var testSlice2 = make([]int, 0, n)
 fmt.Printf("Total time without preallocation: %v \n", timeloop(testSlice, n))
 fmt.Printf("Total time with preallocation: %v \n", timeloop(testSlice2, n))
 //Total time without preallocation: 10.175459ms 
 //Total time with preallocation: 2.22075ms 

}


func timeloop(slice []int, n int) time.Duration {
	 var t0 = time.Now()
	 for len(slice)<n{
		 slice = append(slice, 1)
	 }
	 return time.Since(t0)
}
