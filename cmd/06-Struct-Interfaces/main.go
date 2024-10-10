package main

import "fmt"

// struct and interface

type gasEngine struct { 
	mpg uint8
	gallons uint8
	owner string
	int
}

type electricEngine struct { 
	mpkwh uint8
	kwh uint8
}
func (e gasEngine)  milesLeft() uint8 {
	return e.gallons *e.mpg
}

func (e electricEngine) milesLeft() uint8 { 
	return e.kwh*e.mpkwh
}
// can wraper stuct function milesLeft options
type engine interface {
	  milesLeft() uint8
}

 
func canMakeit( e engine, miles uint8) {
	  if miles<= e.milesLeft() {
		fmt.Println("You can make it there") 
	  } else {
		fmt.Println("Nedd to fuel up first!")
	  }
}

func main() {
	var myEngine gasEngine = gasEngine{25,15, "alex", 10}
	var myEngine2 electricEngine = electricEngine{25,15}
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.owner, myEngine.int)
 // default value myEngine{mpg: 0, gallons: 0}
 // set value myEngine{mpg: 25, gallons: 15 owner: alex, int: 10}


 fmt.Printf("Total miles left in tank: %v", myEngine.milesLeft())

 canMakeit(myEngine, 50) // you can make it there
 canMakeit(myEngine2, 255) // Need to fuel up first

}