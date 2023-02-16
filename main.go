package main
//when importing multiple packages you lock them in parenthesis and put them on their own line
import (
	"fmt"
	"strings"
	"Go-Practice/helper"
	) 

//different syntax for assigning variables
const conferenceTickets = 50
var conferenceName string = "Go Conference"
//uint will not take negative integers
var remainingTickets uint = 50
var bookings = []string{}

func main(){

	greetUsers()

	//for loops have a default true conditional that can be altered to fit your own conditionals
	for remainingTickets > 0 && len(bookings) < 50 {
		//access user input and values
		firstName, lastName, email, userTickets := getUserInput()

		//returned variables pulled from the validateUserInput function
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidEmail && isValidName && isValidTicketNumber {
			//logic for booking ticket to decrement ticket total and who is booking the ticket
			bookTicket(userTickets, firstName, lastName, email)

			//call function print first names
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)
			
			//conditional to prevent booking if event is sold out
			if remainingTickets == 0 {
				fmt.Println("Sorry but this event has been sold out.")
				break
			}
		// conditional to prevent user error and validate user input
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

//print user 
func greetUsers() {
	fmt.Printf("Welcome to the %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here!")
}

//when returning you have function parameters and output parameters 
func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

//prompts to retrieve user name, email, and how many tickets they want
func getUserInput() (string, string, string, uint){
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

	return firstName, lastName, email, userTickets
}

//logic for users to book tickets
func bookTicket(userTickets uint, firstName string, lastName string, email string){
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName + " " + lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets.  You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email);
	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)
}
