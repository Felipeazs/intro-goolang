package helper

import "strings"

//reference to another file: no need to add special sintax for files in the same package (main)
/*when using another package, it must be exported and imported in the main file by changing the
first letter of the function to capital letter in both packages */

// multiple value returns
func ValidateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	//remainingTicket is a global variable: no need to pass it as function arguments
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@") //search for a @ character
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
