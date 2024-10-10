package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Channels

// hold data
// thread safe
// listen for data

var MAX_CHICKEN_PRICE float32 = 5
var MAX_TOFU_PRICE float32 = 3
func main( ){ 
	var c = make(chan int)
	go process(c)
	for i:= range c{

		fmt.Println(i)
	}

	// make more real use

	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)
	var websites =[]string{"walmart.com", "costco.com", "wholefoods.com"}
	for i := range websites{
		 go checkChickenPrices(websites[i], chickenChannel)
		 go checkTofuPrices(websites[i], tofuChannel)
	}
	sendMessage(chickenChannel, tofuChannel)

}

func checkChickenPrices(website string, chickenChannel chan string) {
	  for { 
		time.Sleep(time.Second*1)
		var chickenPrice = rand.Float32()*20
		if chickenPrice<=MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	  }
}

func checkTofuPrices(website string, c chan string) {
	  for { 
		time.Sleep(time.Second*1)
		var tofuPrice = rand.Float32()*20
		if tofuPrice<=MAX_TOFU_PRICE {
			c <- website
			break
		}
	  }
}

func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	// if statement with channel
	select{
	      case website := <-chickenChannel:
				fmt.Printf("\n Text Sent : Fund a deal on chicken at %v",website)
		  case website := <-tofuChannel:
		  		fmt.Printf("\n Email Sent : Found a deal on tofu at %v", website)
	}
}


func process(c chan int) { 
	defer close(c) // after loop finish will be close not awaiting
	for i:=0; i<5; i++{
		 c <- i
	}



	 
	
}