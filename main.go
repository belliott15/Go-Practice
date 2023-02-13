package main
//when importing multiple packages you lock them in parenthesis and put them on their own line
import (
	"fmt"
	"strings"
	) 

func main(){
	//different syntax for assigning variables
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	//uint will not take negative integers
	var remainingTickets uint = 50
	var bookings = []string{}

	fmt.Printf("Welcome to the %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here!")

	//for loops have a default true conditional that can be altered to fit your own conditionals
	for remainingTickets > 0 && len(bookings) < 50 {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// ask user for information
		fmt.Println("Please enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Please enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Please enter your email address: ")
		fmt.Scan(&email)
		
		fmt.Println("Please enter the number of tickets desired ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2 
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidEmail && isValidName && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName + " " + lastName)


			fmt.Printf("Thank you %v %v for booking %v tickets.  You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email);
			fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The first names of all our bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Sorry but this event has been sold out.")
				break
			}
		} else {
			if !isValidName {
				fmt.Printf("The first or last name you entered is not greater than two characters. \n")
			}

			if !isValidEmail {
				fmt.Printf("The email you entered does not contain @. \n")
			}

			if !isValidTicketNumber {
				fmt.Printf("Number of tickets requested is invalid. \n")
			}
		}
	}	
}
