package main

import "fmt"

func main() {
		conferenceName := "Go Conference"
		const conferenceTickets int = 50
		var remainingTickets uint = 50
		var bookings []string		// var bookings = []string{}, bookings := []string{}		
		
		fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

		fmt.Printf("Welcome to %v booking application!\n", conferenceName)
		fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
		fmt.Println("Get your tickets here to attend.")
		
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// ask user for their name
		// fmt.Scan(&firstName)
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)
		
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
		
		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)
		
		fmt.Println("How many tickets do you want?: ")
		fmt.Scan(&userTickets)
		
		remainingTickets = remainingTickets - userTickets		
		bookings = append(bookings, firstName + " " + lastName)

		fmt.Printf("The whole slice: %v\n", bookings)
		fmt.Printf("The first value: %v\n", bookings[0])
		fmt.Printf("Slice type: %T\n", bookings)
		fmt.Printf("Slice size: %v\n", len(bookings))


		fmt.Printf("Thank you %v %v for booking %v tickets!\nYou will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}