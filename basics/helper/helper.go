package main

import "strings"

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {

	// input validation
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@") // gives back boolean result
	isValidTicketNum := userTickets > 0 && userTickets < remainTickets

	return isValidName, isValidEmail, isValidTicketNum
}
