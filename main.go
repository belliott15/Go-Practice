package main
import "fmt" 

func main(){
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to the %v booking application!\n", conferenceName);
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here!");

	var firstName string
	var lastName string
	var email string
	var userTickets int
	// ask user for name
	fmt.Println("Please enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email address: ")
	fmt.Scan(&email)
	
	fmt.Println("Please enter the number of tickets desired ")
	fmt.Scan(&userTickets)

	 fmt.Printf("Thank you %v %v for booking %v tickets.  You will receieve a confirmation email at %v\n", firstName, lastName, userTickets, email);
}
