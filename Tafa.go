package main

import (
	"fmt"
)

func main() {
	var conferenceName = " Tafa Business Conference"
	const conferenceTickets = 30
	var remainingTickets uint = 30
	var bookings []string

	fmt.Println("Jambo, Welcome to", conferenceName, "application")
	fmt.Printf("The %v , where we Build and scale\n", conferenceName)
	fmt.Println("we have a total of", conferenceTickets, "tickets and", remainingTickets, "reamaining tickets")
	fmt.Println(" Get your Tickets here ")

	for {
		var userName string
		var userTickets uint
		var phoneNumber int
		// deatails entry
		fmt.Println("Enter your Name")
		fmt.Scan(&userName)
		fmt.Println("Enter your phone number")
		fmt.Scan(&phoneNumber)

		fmt.Println("Enter Number of tickets to book")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets

		bookings = append(bookings, userName)

		fmt.Printf("The whole Slice ; %v\n", bookings)
		fmt.Printf("First User: %v\n", bookings[0])

		fmt.Printf("Thank you %v  for Purchasing %v Tickets\n", userName, userTickets)

		fmt.Printf("Remaining tickets = %v for the conference\n", remainingTickets)
		fmt.Printf("These are all the bookings made: %v\n", bookings)

		if remainingTickets == 0 {
			fmt.Println("We are fully booked")
			break

		}

	}
}
