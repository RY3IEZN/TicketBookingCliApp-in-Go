package main

import (
	"fmt"
	"strconv"
)

// Go can automatically infer values, or you can use the var,int,uint etc to declare them, me is using :=
// also if a variable is not used, it will be underlined as error... you can tempoarily ignore till used
// although you cant use the := outside the func bloc

const conferenceTickets = 50

var conferenceName string = "Uneku's Conference"
var remainingTickets int = 50

// var booking []string //an array of size-N
// var booking = make([]map[string]string, 0) //an array of size-N
var booking = make([]userData, 0)

type userData struct {
	userName    string
	lastname    string
	userTickets int
}

func main() {

	// Welcome the user,
	// printf can be used just that the variables,will be at the end and in a particular order
	// later we will pass it in a func
	greetUser()

	// use a for loop to keep the app running
	for {

		// get user input
		// was later convereted to func
		userName, lastName, userTickets := getUserInput()

		// validations of the user input
		// was later turn into a func
		isValidLastName, isValidUserName, IsValidTicketRequest := isValid(userName, lastName, userTickets)

		if isValidLastName && isValidUserName && IsValidTicketRequest {

			bookTicket(userTickets, userName, lastName)

			// print only the firstname using a foreach loop
			// step1 create an empy array/slice
			// again, it was moved into a function
			firstName := getFirstNames()
			fmt.Println("the firstnames are:", firstName)

			// check if tickets are remain, if finished end the app
			if remainingTickets == 0 {
				fmt.Print("tickets are sold out")
				break
			}

		} else {
			if !isValidUserName {
				fmt.Println("username is too short")
			}
			if !isValidLastName {
				fmt.Println("lastName is too short")
			}
			if !IsValidTicketRequest {
				fmt.Println("invalid ticket slection")
			}
			println("invalid input")
		}

	}
}

// greet the user
func greetUser() {
	fmt.Println("welcome to the show")
	fmt.Printf("Welcome to the %v booking application \n", conferenceName)
	fmt.Println("Welcome to the", conferenceName, "booking application")
	fmt.Println("We have a total of", conferenceTickets, "and have", remainingTickets, "left")
	fmt.Println("Get your tickets here to attend")
}

// return the firstnames only
func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range booking {
		// step2 slices the values in an existing array
		// var names = strings.Fields(booking)
		// append it into the newly created array
		firstNames = append(firstNames, booking.userName)
	}
	return firstNames
}

// get the user input
func getUserInput() (string, string, int) {

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

	return userName, lastName, userTickets
}

func bookTicket(userTickets int, userName string, lastName string) {
	//remaining tickets
	remainingTickets = remainingTickets - userTickets

	// using a strcut
	var userDatas = userData{
		userName:    userName,
		lastname:    lastName,
		userTickets: userTickets,
	}

	// creating a map
	// import map at the top
	var userData = make(map[string]string)
	userData["userName"] = userName
	userData["lastName"] = lastName
	userData["noTickets"] = strconv.FormatInt(int64(userTickets), 10)

	// dynamically add users to a list
	booking = append(booking, userDatas)
	fmt.Println("-------", booking, "--------")

	// thanks and display remaning ticket
	fmt.Println("thanks, there are", remainingTickets, " tickets left")
	// print the booking list
	fmt.Println("the whole array", booking)

}

func isValid(userName string, lastName string, userTickets int) (bool, bool, bool) {
	isValidUserName := len(userName) >= 2
	isValidLastName := len(lastName) >= 2
	IsValidTicketRequest := userTickets > 0 && userTickets <= remainingTickets
	return isValidLastName, isValidUserName, IsValidTicketRequest

}
