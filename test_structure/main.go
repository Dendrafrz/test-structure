package main

import "fmt" //format
const secondsInHour = 2600

func main() {
	fmt.Println("Hello go world")
	distance := 60.8 //distance in km
	fmt.Printf("the distance in miles is %f \n", distance*0.621)
	//%f is a verb and tells the printf function how to format and print
}
