package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

type User struct {
	firstName string
	lastName  string
	email     string
	tickets   uint
}

var waitGroup = sync.WaitGroup{}

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 30
	var remainingTickets uint = conferenceTickets
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("Total of %v tickets\n", conferenceTickets)
	fmt.Printf("Remaining of %v tickets\n", remainingTickets)

	var bookings = make([]User, 0)

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicket := helper.ValidateInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicket {
			remainingTickets -= userTickets

			var user = User{
				firstName: firstName,
				lastName:  lastName,
				email:     email,
				tickets:   userTickets,
			}

			bookings = append(bookings, user)

			fmt.Printf("Thank you %v %v for booking %v tickets.\nYou will receive a confirmation email at %v\n", user.firstName, user.lastName, user.tickets, user.email)

			waitGroup.Add(1)
			go sendTicket(user)
			userNames := getUserNames(bookings)

			fmt.Printf("Current Bookings: %v\n", userNames)
			fmt.Printf("%v Remaining tickets for %v\n", remainingTickets, conferenceName)

			if remainingTickets == 0 {
				sendTicket(user)
				fmt.Println("The tickets for this Conference have been booked out.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Invalid first name or last name. Names should be longer than 2 characters, please try again")
				continue
			}
			if !isValidEmail {
				fmt.Println("Invalid email, please try again")
				continue
			}
			if !isValidTicket {
				fmt.Printf("There are only %v remaining tickets, please try again\n", remainingTickets)
				continue
			}
		}
		waitGroup.Wait()
	}
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func getUserNames(bookings []User) []string {
	userNames := []string{}
	for _, booking := range bookings {
		userNames = append(userNames, booking.firstName)
	}
	return userNames
}

func sendTicket(user User) {
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", user.tickets, user.firstName, user.lastName)
	fmt.Println("//////////////////")
	fmt.Printf("Sending confirmation details:\n%v\nto your email adress: %v\n", ticket, user.email)
	fmt.Println("//////////////////")
	waitGroup.Done()
}
