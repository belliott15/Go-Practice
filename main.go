package main

//when importing multiple packages you lock them in parenthesis and put them on their own line
import (
	"fmt"
	"strings"
	"time"
)

//different syntax for assigning variables
const conferenceTickets = 50
var conferenceName string = "Go Conference"
//uint will not take negative integers
var remainingTickets uint = 50
//initialize a list of maps
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName string
	email string
	ticketsReserved uint
}

func main(){

	greetUsers()

	//for loops have a default true conditional that can be altered to fit your own conditionals
	for remainingTickets > 0 && len(bookings) < 50 {
		//access user input and values
		firstName, lastName, email, userTickets := getUserInput()

		//returned variables pulled from the validateUserInput function
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidEmail && isValidName && isValidTicketNumber {
			//logic for booking ticket to decrement ticket total and who is booking the ticket
			bookTicket(userTickets, firstName, lastName, email)

			//send email of ticket to user
			//using go keyword it will create a breakout thread to allow the rest of the code to continually run
			//allows concurrent users to book tickets and still remains responsive
			go sendTicket(userTickets, firstName, lastName, email)

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

//validates user input to prevent negative numbers, false emails, and over booking of tickets
func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool){
	isValidName := len(firstName) >= 2 && len(lastName) >= 2 
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

//when returning you have function parameters and output parameters 
func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
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

	//create map for user info
	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		ticketsReserved: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("list of bookings is %v \n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets.  You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email);
	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#####################")
	fmt.Printf("Sending ticket: \n %v  \n to email address %v \n", ticket, email)
	fmt.Println("#####################")
}
