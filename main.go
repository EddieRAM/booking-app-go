package main

import (
	"booking-app/validations"
	"fmt"
	"strings"
)

// package level variables

const conferenceTickets int = 50 
var conferenceName = "Go Conference" // sintax requires var to be accessible at the package level
var remainingTickets uint = 50
var bookings []string		// var bookings = []string{}, bookings := []string{}

func main() {
				
		
		greetUsers() // conferenceName, conferenceTickets, remainingTickets comes from the package level

		fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
				
		for {
			
					firstName, lastName, email, userTickets := getUserInputs()

					isValidName, isValidEmail, isValidTicketNumber := validations.ValidateInputs(firstName, lastName, email, userTickets, remainingTickets)
					
					if isValidName && isValidEmail && isValidTicketNumber {
						
						createBookings(userTickets, firstName, lastName, email)						

						firstNames := pickFirstNames()
						fmt.Printf("The first names of bookings are: %v\n", firstNames)
						
						// To show only the first names in the slice						
						
						if remainingTickets == 0 {
							fmt.Println("Our conference is booked out! Come back next year.")
							// end program
								break
						}
					} else {
							// validations
							if !isValidName {
								fmt.Println("The first name or last name you entered is too short")
							}
							if !isValidEmail {
								fmt.Println("The email address you entered does'nt contain @ sign")
							}
							if !isValidTicketNumber {
								fmt.Println("The Number of tickets you entered is invalid")
							}										
					}
			}
}

func greetUsers() {
		fmt.Printf("Welcome to %v booking application\n", conferenceName)	// names of the variables must be equal to the defined ones at the package level
		fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
		fmt.Println("Get your tickets here to attend.")
}

func pickFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings { // index has been replaced with _, because it has never been used 
			var names = strings.Fields(booking)	// strings.Fields => slice manipulation method					
			firstNames = append(firstNames, names[0]) // appends firstNames slice with firstName
	}
	return firstNames
}

func getUserInputs() (string, string, string, uint){
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

		return firstName, lastName, email, userTickets
}

func createBookings(userTickets uint, firstName string, lastName string, email string) {
		remainingTickets = remainingTickets - userTickets		
		bookings = append(bookings, firstName + " " + lastName)
		
		fmt.Printf("Thank you %v %v for booking %v tickets!\nYou will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)		
}