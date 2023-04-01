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

	// greetUser(conferenceName)

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

		//input validation before running the app
		isNameValid := len(firstName) >= 2 && len(lastName) >= 2
		isEmailValid := strings.Contains(email, "@")
		isTicketAvailable := userTicket > 0 && userTicket <= remainingTickets

		//booking logic, tickets left,saving details

		// validate then run or break
		if isNameValid && isEmailValid && isTicketAvailable {

			//get thr number of remaining ticket
			remainingTickets = remainingTickets - userTicket
			//aray for the ppl
			bookings = append(bookings, firstName+" "+lastName)

			//thank the customer for buying, say how many tickets left
			fmt.Println("thank you", firstName, lastName, "for buying", userTicket, "ticket, you will get a confirmation email in ", email)
			fmt.Println("there are now ", remainingTickets, " tickets left")
			fmt.Println("here are all the bookings made ", bookings)

			//gets thr firstnames
			var firstNames = getFirstNames(bookings)
			fmt.Println("theses are the first names of bookings", firstNames)

			//condition 1 to break the loop,
			//{if the tickets have finished, inform them that it has finished
			// and stop the service }
			if remainingTickets == 0 {
				fmt.Print("All tickets are sold out, come back next week")
				//stop the loop
				break
			}
		} else {
			if !isNameValid {
				fmt.Println(" firstname or lastname is too short,Please input 3 or more characters")
			}
			if !isEmailValid {
				fmt.Println("the email address is invalid and doesnt contain @")
			}
			if !isTicketAvailable {
				fmt.Println("invalid number of tickets selection")
			}
			continue
		}
	}
}

func getFirstNames(bookings []string) []string {
	//print only the 1st name
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}
