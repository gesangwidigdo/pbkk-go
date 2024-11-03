package main

import (
	"fmt"
	"pbkk-go/helper"
	"sync"
	"time"
)

// Variable and constant
var conferenceName = "Go Conference" // should manually use var when declaring variable in package level
const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName   string
	lastName    string
	email       string
	userTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	firstName, lastName, email, userTickets := getUserInput()

	// Call validate function
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		// Call book ticket function
		bookTicket(firstName, lastName, email, userTickets)

		wg.Add(1)
		go sendTicket(firstName, lastName, email, userTickets)

		// Call print first names function
		firstNames := getFirstNames()
		fmt.Printf("These are our booking list: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("All tickets are sold out.")
			// break
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
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{} // -> local variable, only can be used in this scope
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTicket(firstName string, lastName string, email string, userTickets uint) {
	remainingTickets = remainingTickets - userTickets

	// create map for user
	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		userTickets: userTickets,
	}
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	// storing to slices
	bookings = append(bookings, userData)
	fmt.Printf("List of booking is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(firstName string, lastName string, email string, userTickets uint) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("Sending %v tickets to %v %v at %v\n", userTickets, firstName, lastName, email)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Printf("Sending ticket:\n%v\nto email address: %v\n", ticket, email)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	wg.Done()
}
