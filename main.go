package main

import "fmt"

func main(){
    var conferenceName = "Go conference"
    const conferenceTickets = 50
    var remainingTickets = 50
    
    fmt.Printf("Welcome to our %v booking application\n", conferenceName)
    fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
    fmt.Println("Get your tickets here to attend")


    var userName string
    var userTickets int
    //ask user

    userName = "Felipe"
    userTickets = 2

    fmt.Printf("userName is %T, userTickets are %T", userName, userTickets)
    fmt.Printfn("User %v booked %v tickets.", userName, userTickets)

}






