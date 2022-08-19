package main

import (
	"booking-app/validations"
	"fmt"
)

// package level variables

const conferenceTickets int = 50 
var conferenceName = "Go Conference" // sintax requires var to be accessible at the package level
var remainingTickets uint = 50
var bookings = make([]UserData, 0)	// now bookings uses the struct UserData

// create a type struct to get store information with multiple datatypes

type UserData struct {
		firstName string
		lastName string
		email string
		numberOftickets uint
}

func main() {				
		
		greetUsers() // conferenceName, conferenceTickets, remainingTickets comes from the package level

		fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
				
		for {
			
					firstName, lastName, email, userTickets := getUserInputs()

					isValidName, isValidEmail, isValidTicketNumber := validations.ValidateInputs(firstName, lastName, email, userTickets, remainingTickets)
					
					if isValidName && isValidEmail && isValidTicketNumber {
						
						createBookings(userTickets, firstName, lastName, email)
						sendTicket(userTickets, firstName, lastName, email)					

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
				firstNames = append(firstNames, booking.firstName) // structs use dot notation
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
		
		// assign values to the struct keys
				
		var userData = UserData {
			firstName: firstName,
			lastName: lastName,
			email: email,
			numberOftickets: userTickets,
		}		

		bookings = append(bookings, userData)
		fmt.Printf("List of bookings is %v\n", bookings)
		
		fmt.Printf("Thank you %v %v for booking %v tickets!\nYou will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)		
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	var ticket = fmt.Sprintf("%v ticket(s) for %v %v", userTickets, firstName, lastName)
	fmt.Println("##########################")
	fmt.Printf("Sending ticket\n.\n..\n...\n Congratulations! %v\n to email address %v was sent!\n", ticket, email)
	fmt.Println("##########################")
}