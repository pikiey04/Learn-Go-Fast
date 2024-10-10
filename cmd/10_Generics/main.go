package main

import "fmt"

//generics working with go
// receive additional params and type in func

 type gasEngine struct {
	gallons float32
	mpg float32
 }

 type electricEngine struct{
	kwh float32
	mpkwh float32
 }

 // use generics type struct [T type,type]
 type car [T gasEngine | electricEngine]struct {
	carMake string
	carModel string
	engine T
}


 func main( ) { 
  var gasCar = car[gasEngine] {
	 carMake: "Honda",
	 carModel: "Civice",
	 engine: gasEngine{
		gallons: 12.4,
		 mpg: 40,
	 },
  }
  var electricCar = car[electricEngine] {
	 carMake: "Tesla",
	 carModel: "Model 3",
	 engine: electricEngine{
		kwh: 57.5,
		mpkwh: 4.17,
	 },
  }
   fmt.Println(gasCar)
   fmt.Println(electricCar)
   
 }