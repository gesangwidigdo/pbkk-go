package main

import (
	"fmt"
	"strings"
)

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

	// // Data Types
	// var firstName string
	// var lastName string
	// var email string
	// var userTickets uint

	// // declaring array with value
	// // var bookings = [50]string{"John Doe", "Jane Doe", "Alice Doe"} --> fixed size array

	// // array declaration
	// // var bookings [50]string

	// // slice declaration
	// // var bookings []string
	// // alternatives:
	// bookings := []string{}

	// // ask user for input
	// fmt.Print("Enter your first name: ")
	// fmt.Scan(&firstName)
	
	// fmt.Print("Enter your last name: ")
	// fmt.Scan(&lastName)

	// fmt.Print("Enter your email address: ")
	// fmt.Scan(&email)

	// fmt.Print("Enter number of tickets: ")
	// fmt.Scan(&userTickets)

	// remainingTickets = remainingTickets - userTickets
	
	// // storing to array
	// // bookings[0] = firstName + " " + lastName

	// // storing to slices
	// bookings = append(bookings, firstName + " " + lastName)

	// // fmt.Printf("The whole array/slice: %v\n", bookings)
	// // fmt.Printf("The first val of array/slice: %v\n", bookings[0])
	// // fmt.Printf("The type of array/slice: %T\n", bookings)
	// // fmt.Printf("The length of array/slice: %v\n", len(bookings))

	// fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	// fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	// fmt.Printf("These are our booking list: %v\n", bookings)

	bookings := []string{}
	for {
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

		isValidName := len(firstName) > 2 && len(lastName) > 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets

			// storing to slices
			bookings = append(bookings, firstName + " " + lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("These are our booking list: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("All tickets are sold out.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name is too short.")
			}
			if !isValidEmail {
				fmt.Println("email is invalid.")
			}
			if !isValidTicketNumber {
				fmt.Printf("Sorry, we only have %v tickets remaining.\n", remainingTickets)
			}
		}
	}

	city := "Jakarta"

	switch city {
	case "New York":
		fmt.Println("Conference location is New York")
	
	case "Jakarta", "Singapore", "Hong Kong", "Tokyo":
		fmt.Println("Conference location is Asia")

	case "London", "Paris", "Berlin":
		fmt.Println("Conference location is Europe")
	
	default:
		fmt.Println("Conference location is unknown")
	}
}