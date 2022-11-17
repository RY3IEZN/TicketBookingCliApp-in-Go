package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Nicks conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	bookings := []string{}

	// welcome message
	fmt.Println("Welcome to", conferenceName)
	fmt.Println("there are only", conferenceTickets, "and they are", remainingTickets, "tickets left")

	//continue running the app until the loop is broken
	for {

		// getting the customer details
		var firstName string
		var lastName string
		var email string
		var userTicket int

		// ask for firstName
		fmt.Println("Enter your First name: ")
		//get the input value
		fmt.Scan(&firstName)

		// ask for lastName
		fmt.Println("Enter your Last name: ")
		//get the input value
		fmt.Scan(&lastName)

		// ask for email
		fmt.Println("Enter your email: ")
		//get the input value
		fmt.Scan(&email)

		// ask for no. of tickets to buy
		fmt.Println("How many tickets do you want to buy: ")
		//get the input value
		fmt.Scan(&userTicket)

		//booking logic, tickets left,saving details
		//
		remainingTickets = remainingTickets - userTicket
		//aray for the ppl
		// bookings[0] = firstName + lastName
		bookings = append(bookings, firstName+" "+lastName)

		//thank the customer for buying, say how many tickets left
		fmt.Println("thank you", firstName, lastName, "for buying", userTicket, "ticket, you will get a confirmation email in ", email)
		fmt.Println("there are now ", remainingTickets, " tickets left")
		fmt.Println("here are all the bookings made ", bookings)

		//print only the 1st name
		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Println("theses are the first names of bookings", firstNames)

		//condition 1 to break the loop,
		//{if the tickets have finished, inform them that it has finished
		// and stop the service }
		if remainingTickets == 0 {
			fmt.Print("All tickets are sold out, come back next week")
			//stop the loop
			break
		}
	}
}
