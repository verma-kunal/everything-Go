package main

import (
	"Go-Learn/basics/helper"
	"fmt"
	"sync"
	"time"
)

// package-level variables:
var confName = "GopherCon Community Edition 2024"

const confTickets = 50      // total tickets available
var remainTickets uint = 50 // this value reduces everytime someone buys a ticket
// var bookings = make([]map[string]string, 0) // booking - a list of maps
var bookings = make([]UserData, 0) // booking - a list of UserData struct

type UserData struct {
	firstName    string
	lastName     string
	email        string
	numOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	// looping the entire process:
	for {

		// function for taking user input:
		firstName, lastName, email, userTickets := userInput()

		// function for user input validation:
		isValidName, isValidEmail, isValidTicketNum := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainTickets)

		// check
		if isValidName && isValidEmail && isValidTicketNum { // all 3 input validations have to be true to execute the loop

			// function containing booking ticket logic:
			remainingTickets := bookingTickets(remainTickets, userTickets, firstName, lastName, email)
			fmt.Printf("The no. of remaining tickets are: %v\n", remainingTickets)

			// calling the sendTicket function:
			wg.Add(1) // num of goroutines/thread to wait for, before creating a new thread
			go sendTicket(userTickets, firstName, lastName, email)

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
	wg.Wait() // waits for all threads to be executed
}

func bookingTickets(remainTickets uint, userTickets uint, firstName string, lastName string, email string) uint {
	// book tickets logic:
	remainTickets = remainTickets - userTickets

	// populating the UserData struct with the actual values:
	userData := UserData{
		firstName:    firstName,
		lastName:     lastName,
		email:        email,
		numOfTickets: userTickets,
	}

	// using a slice of map & appending a new value to it:
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is: %v\n", bookings)

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

func greetUsers() {
	fmt.Printf("Welcome to %v booking application!", confName)
	fmt.Printf("We have total of %v tickets and %v are still remaining to be purchased!\n", confTickets, remainTickets)
	fmt.Println("!!!! Get your tickets here !!!!")
}
func getFirstNames() []string {

	// looping through our bookings slice & grabbing only the 'first names'
	var firstNames []string
	for _, booking := range bookings {
		// adding the firstname to our new slice
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

// sendTicket generates a ticket & send via email
func sendTicket(userTickets uint, firstName string, lastName string, email string) {

	// delay the execution:
	time.Sleep(50 * time.Second)

	ticket := fmt.Sprintf("%v tickets are for: %v %v/n", userTickets, firstName, lastName)
	fmt.Println("##################")
	fmt.Printf("Sending ticket:\n %v \nto email %v\n", ticket, email)
	fmt.Println("##################")

	wg.Done() // removes the added goroutine from the WaitGroup - when the execution is done!
}
