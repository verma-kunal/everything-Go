package main

import "fmt"

func main() {

	confName := "GopherCon Community Edition 2024"
	const confTickets = 50      // total tickets available
	var remainTickets uint = 50 // this value reduces everytime someone buys a ticket
	var bookings []string       // booking info slice (variable-length)

	fmt.Printf("Welcome to %v !\n", confName)
	fmt.Printf("We have total of %v tickets and %v are still remaining to be purchased!\n", confTickets, remainTickets)
	fmt.Println("!!!! Get your tickets here !!!!")

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

	// book tickets logic:
	remainTickets = remainTickets - userTickets

	// using a slice & appending a new value to it:
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you '%v %v' for booking '%v' tickets today! You'll receive a confirmation at email: '%v'\n", firstName, lastName, userTickets, email)
	fmt.Printf("Number of ticket remaining now: %v", remainTickets)

	fmt.Printf("The total bookings: %v\n", bookings)

}
