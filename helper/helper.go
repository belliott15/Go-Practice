package helper

import (
	"strings"
)

//the way to export functions to another file is to capitalize the first letter of the function name

//validates user input to prevent negative numbers, false emails, and over booking of tickets
func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool){
	isValidName := len(firstName) >= 2 && len(lastName) >= 2 
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

