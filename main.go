package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go conference" //type inference :=
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	bookings := []string{} //list of strings

	//for loop
	for {

		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Printf("firstName is %T\n", firstName)
		fmt.Printf("userTickets is %T\n", userTickets)

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName) //print the value in memory -- & is a pointer

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName) //print the value in memory -- & is a pointer

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		bookings = append(bookings, firstName+" "+lastName)

		remainingTickets = remainingTickets - userTickets

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n",
			firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining", remainingTickets)

		firstNames := []string{}
		//list iteration for each
		//blank identifier (_) replaces a variable for nothing; in this case repalaces the index because it is not used
		for _, booking := range bookings {
			var names = strings.Fields(booking) // splits the string with empty spaces as separator "Felipe Zapata" => ["Felipe", "Zapata"]
			var firstName = names[0]
			firstNames = append(firstNames, firstName)
		}

		fmt.Printf("These are all our bookings: %v\n", firstNames)
	}
}
