package main

import "fmt"

func main() {
	// Variable and constant

	conferenceName := "Go Conference" // syntactic sugar, can't used for defining constant
	const conferenceTickets = 50
	var remainingTickets uint = 50

	// remainingTickets = -10 // error, bcs it's unsigned integer

	// Check type of variables
	// fmt.Printf("conferenceName: %T, conferenceTickets: %T, remainingTickets: %T\n", conferenceName, conferenceTickets, remainingTickets)
	
	// Formatted output
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
	
	// conferenceTickets = 30 --> error

	// Data Types
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask user for input
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)
	
	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}