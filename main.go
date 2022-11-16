package main

import "fmt"

func main() {
	conferenceName := "Nicks conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Welcome to", conferenceName)
	fmt.Println("there are only", conferenceTickets, "and they are", remainingTickets, "tickets left")

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

	// ask for no of tickets to buy
	fmt.Println("How many tickets do you want to buy: ")
	//get the input value
	fmt.Scan(&userTicket)

	remainingTickets = remainingTickets - userTicket

	fmt.Println("thank you", firstName, lastName, "for buying", userTicket, "ticket, you will get a confirmation email in your email", email)

	fmt.Print("there are now", remainingTickets, "tickets left")
}
