package main

import (
	"fmt"
	"strings"
)

func main() {
	// Go can automatically infer values, or you can use the var,int,uint etc to declare them, me is using :=
	// also if a variable is not used, it will be underlined as error... you can tempoarily ignore till used
	conferenceName := "Uneku's Conference"
	const conferenceTickets = 50
	remainingTickets := 50
	var booking []string //an array of size-N

	// Welcome the user,
	// printf can be used just that the variables,will be at the end and in a particular order
	fmt.Printf("Welcome to the %v booking application \n", conferenceName)
	fmt.Println("Welcome to the", conferenceName, "booking application")
	fmt.Println("We have a total of", conferenceTickets, "and have", remainingTickets, "left")
	fmt.Println("Get your tickets here to attend")

	// use a for loop to keep the app running
	for {

		// data types, when you declare variable, you will need to specify the type, int,str,bool
		// ask for the user's name and number of tickets thet want and then assign it to the variable using pointers
		// arrays require a size, the 3 arrays shown below are the same thing, just expressed easier
		var userName string
		var lastName string
		var userTickets int
		// var booking = [50]string{}
		// var bookingReg [50]string
		// booking := []string{}

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

		if userTickets <= remainingTickets {

			//remaining tickets
			remainingTickets = remainingTickets - userTickets

			// dynamically add users to a list
			booking = append(booking, userName+" "+lastName)

			// thanks and display remaning ticket
			fmt.Println("thanks, there are", remainingTickets, " tickets left")
			// print the booking list
			fmt.Println("the whole array", booking)

			// print only the firstname using a foreach loop
			// step1 create an empy array/slice
			firstNames := []string{}
			for _, booking := range booking {
				// step2 slices the values in an existing array
				var names = strings.Fields(booking)
				// append it into the newly created array
				firstNames = append(firstNames, names[0])
			}
			fmt.Println("the firstnames are:", firstNames)

			// check if tickets are remain, if finished end the app
			if remainingTickets == 0 {
				fmt.Print("tickets are sold out")
				break
			}

		} else {
			println("we only have", remainingTickets, "left", "you cant book that ")
		}

	}
}
