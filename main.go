package main

import "fmt"

func main(){
    conferenceName := "Go conference"
    const conferenceTickets = 50
    var remainingTickets = 50
    
    fmt.Printf("Welcome to our %v booking application\n", conferenceName)
    fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
    fmt.Println("Get your tickets here to attend")


    var firstName string
    var lastName string
    var email string
    var userTickets int

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

    fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
    
    remainingTickets = remainingTickets - userTickets
    fmt.Printf("%v tickets remaining", remainingTickets)

}






