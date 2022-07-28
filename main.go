package main

import (
	"fmt"
	"strings"
)

func main() {
		conferenceName := "Go Conference"
		const conferenceTickets int = 50
		var remainingTickets uint = 50
		var bookings []string		// var bookings = []string{}, bookings := []string{}		
		
		fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
		
		fmt.Printf("Welcome to %v booking application!\n", conferenceName)
		fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
		fmt.Println("Get your tickets here to attend.")
		
		
		for {
				var firstName string
				var lastName string
				var email string
				var userTickets uint
				
				fmt.Println("Enter your first name: ")
				fmt.Scan(&firstName)
				
					fmt.Println("Enter your last name: ")
					fmt.Scan(&lastName)
					
					fmt.Println("Enter your email address: ")
					fmt.Scan(&email)
					
					fmt.Println("How many tickets do you want?: ")
					fmt.Scan(&userTickets)
					
					if userTickets > remainingTickets {
							fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
							continue // restarts the loop
					}

					remainingTickets = remainingTickets - userTickets		
					bookings = append(bookings, firstName + " " + lastName)
					
					fmt.Printf("Thank you %v %v for booking %v tickets!\nYou will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
					fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
					
					// To show only the first names in the slice

					firstNames := []string{}
					for _, booking := range bookings { // index has been replaced with _, because it has never been used 
							var names = strings.Fields(booking)						
							firstNames = append(firstNames, names[0]) // appends firstNames slice with firstName
					}

					fmt.Printf("The first names of bookings are: %v\n", firstNames)
					
					if remainingTickets == 0 {
							// end program
							fmt.Println("Our conference is booked out! Come back next year.")
							break
					}
		}
}