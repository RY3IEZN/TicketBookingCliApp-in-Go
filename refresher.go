package main

import "fmt"

func main() {
	// Go can automatically infer values, or you can use the var,int,uint etc to declare them, me is using :=
	// also if a variable is not used, it will be underlined as error... you can tempoarily ignore till used
	conferenceName := "Uneku's Conference"
	const conferenceTickets = 50
	remainingTickets := 50

	// Welcome the user,
	// printf can be used just that the variables,will be at the end and in a particular order
	fmt.Printf("Welcome to the %v booking application \n", conferenceName)
	fmt.Println("Welcome to the", conferenceName, "booking application")
	fmt.Println("We have a total of", conferenceTickets, "and have", remainingTickets, "left")
	fmt.Println("Get your tickets here to attend")

	// data types, when you declare variable, you will need to specify the type, int,str,bool
	// ask for the user's name and number of tickets thet want and then assign it to the variable using pointers
	// arrays require a size, the 2 arrays shown below are the same thing, just expressed easier
	var userName string
	var lastName string
	var userTickets int
	var booking = [50]string{}
	var bookingReg [50]string

	//pointer is used to set the value to a variable, fmt.scan is to get the user input
	// get the 1st name
	fmt.Printf("Enter your firstname: ")
	fmt.Scan(&userName)
	fmt.Println(userName)

	// get the lastname
	fmt.Printf("Enter your lastname: ")
	fmt.Scan(&lastName)
	fmt.Println(lastName)

	// get number of tickets
	fmt.Printf("Enter the number of tickets you want: ")
	fmt.Scan(&userTickets)
	fmt.Println(userTickets)

	//remaining tickets
	remainingTickets = remainingTickets - userTickets
	println("thanks, there are", remainingTickets, " tickets left")
}
