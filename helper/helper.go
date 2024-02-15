package helper

import "regexp"

func ValidateInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	emailRegex, _ := regexp.Compile("[^@ \t\r\n]+@[^@ \t\r\n]+\\.[^@ \t\r\n]+")
	isValidEmail := emailRegex.MatchString(email)
	isValidTicket := userTickets <= remainingTickets && userTickets > 0
	return isValidName, isValidEmail, isValidTicket
}
