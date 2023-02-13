package main
import "fmt" 

func main(){
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings = []string{}

	fmt.Printf("Welcome to the %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here!")

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for name
	fmt.Println("Please enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email address: ")
	fmt.Scan(&email)
	
	fmt.Println("Please enter the number of tickets desired ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName + " " + lastName)

	fmt.Printf("The whole Slice: %v\n", bookings)
	fmt.Printf("The first Value: %v\n", bookings[0])
	fmt.Printf("The slice type: %T\n", bookings)
	fmt.Printf("The slice length: %v\n", len(bookings))

	 fmt.Printf("Thank you %v %v for booking %v tickets.  You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email);
	 fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)
}
