package main

import (
	"fmt"
)

type userData struct {
	userName    string
	phoneNumber string
	userTickets uint
}

func createUserData(name string, phone string, tickets uint) *userData {
	return &userData{
		userName:    name,
		phoneNumber: phone,
		userTickets: tickets,
	}
}
func bookingProcess(remainingTickets *uint, userTickets uint) {
	*remainingTickets = *remainingTickets - userTickets

}
func main() {
	var conferenceName = " Tafa Business Conference"
	const conferenceTickets = 30
	var remainingTickets uint = 30
	var bookings []*userData

	fmt.Println("Jambo, Welcome to", conferenceName, "application")
	fmt.Printf("The %v , where we Build and scale\n", conferenceName)
	fmt.Println("we have a total of", conferenceTickets, "tickets and", remainingTickets, "reamaining tickets")
	fmt.Println(" Get your Tickets here ")

	for {
		var userName string
		var userTickets uint
		var phoneNumber string

		// deatails entry

		fmt.Println("Enter your Name")
		fmt.Scan(&userName)
		fmt.Println("Enter your phone number(e.g. 0712345678):)")
		fmt.Scan(&phoneNumber)

		if len(phoneNumber) < 10 {
			fmt.Println("Invalid phone number, Numbers must be atleast 10 long")
			continue

		}

		fmt.Println("Enter Number of tickets to book")
		fmt.Scan(&userTickets)

		if userTickets > remainingTickets {
			fmt.Printf("Sorry but, we only have %v Tickets left\n", remainingTickets)
			continue

		}

		bookingProcess(&remainingTickets, userTickets)

		var userPointer *userData = createUserData(userName, phoneNumber, userTickets)

		bookings = append(bookings, userPointer)

		//pointer info check

		fmt.Println("\n--- Booking Summary ---")
		fmt.Printf("Thank you %v for purchasing %v tickets.\n", userPointer.userName, userPointer.userTickets)

		fmt.Printf("Remaining tickets: %v\n", remainingTickets)

		fmt.Println("\n--- All Bookings ---")
		for i, booking := range bookings {
			// Print formatted struct values (%+v prints field names alongside values)
			fmt.Printf("[%d] %v (%v tickets) - Phone: %v\n", i+1, booking.userName, booking.userTickets, booking.phoneNumber)
		}
		fmt.Println("------------------------\n")

		if remainingTickets == 0 {
			fmt.Println("We are fully booked!")
			break
		}
	}
}
