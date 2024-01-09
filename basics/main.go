package main

import (
	"fmt"
	"strings"
)

// package-level variables:
var confName = "GopherCon Community Edition 2024"

const confTickets = 50      // total tickets available
var remainTickets uint = 50 // this value reduces everytime someone buys a ticket
var bookings []string       // booking info slice (variable-length)

func main() {

	greetUsers()

	// looping the entire process:
	for {

		// function for taking user input:
		firstName, lastName, email, userTickets := userInput()

		// function for user input validation:
		isValidName, isValidEmail, isValidTicketNum := validateUserInput(firstName, lastName, email, userTickets)

		// check
		if isValidName && isValidEmail && isValidTicketNum { // all 3 input validations have to be true to execute the loop

			// function containing booking ticket logic:
			remainingTickets := bookingTickets(remainTickets, userTickets, firstName, lastName, email)
			fmt.Printf("The no. of remaining tickets are: %v\n", remainingTickets)

			// function to get the firstnames:
			firstNames := getFirstNames()
			fmt.Printf("The first names of the bookings are: %v\n", firstNames)

			// end the loop if no tickets are left:
			if remainTickets == 0 {
				// end program
				fmt.Println("Tickets are sold out! See you next year.")
				break
			}
		} else if userTickets == remainTickets {
			// just an example to showcase 'else if'

		} else {
			if !isValidName {
				fmt.Println("First or Last name you entered is too short!")
			}
			if !isValidEmail {
				fmt.Println("Email Address entered doesn't contain '@' sign!")
			}
			if !isValidTicketNum {
				fmt.Println("Number of tickets entered is invalid!")
			}
		}
	}
}

func bookingTickets(remainTickets uint, userTickets uint, firstName string, lastName string, email string) uint {
	// book tickets logic:
	remainTickets = remainTickets - userTickets

	// using a slice & appending a new value to it:
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you '%v %v' for booking '%v' tickets today! You'll receive a confirmation at email: '%v'\n", firstName, lastName, userTickets, email)

	return remainTickets
}

func userInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// user input
	fmt.Print("Enter your first name:")
	fmt.Scan(&firstName)
	fmt.Print("Enter your last name:")
	fmt.Scan(&lastName)
	fmt.Print("Enter your email address:")
	fmt.Scan(&email)
	fmt.Print("Enter the number of tickets you wish to buy:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {

	// input validation
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@") // gives back boolean result
	isValidTicketNum := userTickets > 0 && userTickets < remainTickets

	return isValidName, isValidEmail, isValidTicketNum
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application!", confName)
	fmt.Printf("We have total of %v tickets and %v are still remaining to be purchased!\n", confTickets, remainTickets)
	fmt.Println("!!!! Get your tickets here !!!!")
}
func getFirstNames() []string {

	// looping through our bookings slice & grabbing only the 'first names'
	var firstNames []string
	for _, booking := range bookings {
		names := strings.Fields(booking) // converts the string name into a slice (elements separated by ',')

		// adding the firstname to our new slice
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}
