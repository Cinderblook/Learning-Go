package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName   string
	lastName    string
	email       string
	userTickets int
	//isOptedInForNewsLetter bool
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketCount := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketCount {

			bookTicket(userTickets, firstName, lastName, email, conferenceName)

			// Tell main thread to wait before exiting for extra threads to finish
			wg.Add(1)
			// Use Go to have this step operate on another thread
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			// Exit application
			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Printf("First: %v; or last: %v; name entered is invalid.\n", firstName, lastName)
			} else if !isValidEmail {
				fmt.Printf("Email entered: %v; is invalid.\n", email)
			} else if !isValidTicketCount {
				fmt.Printf("Ticket purchase count entered: %v; is invalid.\n", userTickets)
			}
			//fmt.Printf("Your input data is invalid, try again.\n")
		}

	}

}

func greetUsers(conferenceName string, conferenceTickets int, initialRemainingTickets int) {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, initialRemainingTickets)
	fmt.Println("Get your tickets here to attend")
	fmt.Println("------------")

	remainingTickets = initialRemainingTickets
}

func getUserInput() (string, string, string, int) {
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

//	func getFirstNames() []string {
//		// Split strings to display only first names of users buying tickets
//		firstNames := []string{}
//		for _, booking := range bookings {
//			var names = strings.Fields(booking)
//			firstNames = append(firstNames, names[0])
//		}
//		return firstNames
//	}
func getFirstNames() []string {
	// Split strings to display only first names of users buying tickets
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets int, remainingTickets int) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketCount := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketCount
}

func bookTicket(userTickets int, firstName string, lastName string, email string, conferenceName string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		userTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets int, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("####################")
	fmt.Printf("Sending ticket: \n %v \nto email address %v\n", ticket, email)
	fmt.Println("####################")
	// Tell main thread this thread has finished
	wg.Done()
}
