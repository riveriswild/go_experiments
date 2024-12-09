package helper

import "strings"

func ValidateUserInput(userName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	var isValidName bool = len(userName) >= 2 && len(lastName) >=2
	isValidEmail := strings.Contains(email, "@")
	validTicketCount := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, validTicketCount
}