package main

import (
	"fmt"
	"strings"
)

// package level variables, available for main ann all functions
// inference declaration connot be done
const conferenceTickets = 50

var conferenceName = "Go conference" //type inference :=
var remainingTickets uint = 50
var bookings = []string{} //list of strings

func main() {

	greetUsers()

	//for loop
	for { // the same as for true
		firstName, lastName, email, userTickets := getUserInput()

		//input validation
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)

			//call print firstNames
			firstNames := getFirstNames(bookings)
			fmt.Printf("These are all our bookings: %v\n", firstNames)

			//conditionals
			noTicketsRemaining := remainingTickets == 0 //boolan variable by inference, if remainingTickets is cero then noThicketRemainig is true
			if noTicketsRemaining {                     //expression to evaluate
				fmt.Print("Our conference id booked out. Come back next year")
				break //ends the loop
			}
		} else {
			if !isValidName {
				fmt.Printf("first name or lastname you entered is too short\n")
			}

			if !isValidEmail {
				fmt.Printf("emil address you enteres doesn't contain a @ symbol\n")
			}

			if !isValidTicketNumber {
				fmt.Printf("number of tickets you enteres is invalid\n")
			}

			fmt.Printf("your input data is invalid, try again\n")
		}
	}

}

// functions
func greetUsers() {
	//As the variables used here are global, is not necesary to pass them as function arguments
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	//list iteration for each
	//blank identifier (_) replaces a variable for nothing; in this case repalaces the index because it is not used
	//bookings is a global variable: no need to pass it as function arguments
	for _, booking := range bookings {
		var names = strings.Fields(booking) // splits the string with empty spaces as separator "Felipe Zapata" => ["Felipe", "Zapata"]
		var firstName = names[0]
		firstNames = append(firstNames, firstName)
	}
	return firstNames
}

// multiple value returns
func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	//remainingTicket is a global variable: no need to pass it as function arguments
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@") //search for a @ character
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	//checking the variables type (T)
	// fmt.Printf("firstName is %T\n", firstName)
	// fmt.Printf("userTickets is %T\n", userTickets)

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName) //scan print the value in memory -- & is a pointer

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName) //scan print the value in memory -- & is a pointer

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	//remainingTickets, bookings and conferenceName are global variables: no need to pass them as function arguments
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n",
		firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v", remainingTickets, conferenceName)

}
