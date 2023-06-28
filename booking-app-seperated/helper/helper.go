package helper

import (
	"fmt"
	"strings"
)

var remainingTickets int
var bookings []string

func GreetUsers(conferenceName string, conferenceTickets int, initialRemainingTickets int) {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, initialRemainingTickets)
	fmt.Println("Get your tickets here to attend")
	fmt.Println("------------")

	remainingTickets = initialRemainingTickets
}

func GetUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("------------")
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("------------")
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)
	fmt.Println("------------")
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)
	fmt.Println("------------")

	return firstName, lastName, email, userTickets
}

func GetFirstNames() []string {
	// Split strings to display only first names of users buying tickets
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func ValidateUserInput(firstName string, lastName string, email string, userTickets int, remainingTickets int) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketCount := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketCount
}

func BookTicket(userTickets int, firstName string, lastName string, email string, conferenceName string) int {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	return remainingTickets
}
