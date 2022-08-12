package validations

import (
	"strings"
)

// Global variables are always capitalized
func ValidateInputs(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) { // when returning multiple values, the return must be encapsulated
	isValidName := len(firstName) >= 2 && len(lastName) >=2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber // go can return multiple values
}
