package main

import (
	"booking-app-seperated/helper"
	"fmt"
)

func main() {
	const conferenceTickets = 50

	var conferenceName = "Go Conference"
	var remainingTickets = 50
	//var bookings = []string{}
	helper.GreetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {
		firstName, lastName, email, userTickets := helper.GetUserInput()

		isValidName, isValidEmail, isValidTicketCount := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketCount {
			remainingTickets := helper.BookTicket(userTickets, firstName, lastName, email, conferenceName)

			firstNames := helper.GetFirstNames()
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
