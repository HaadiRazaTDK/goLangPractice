package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50

	fmt.Print("----------------------------------------------------\n")
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Print("----------------------------------------------------\n")
	fmt.Printf("We have a total of %v tickets and %v are left.\n", conferenceTickets, remainingTickets)
	fmt.Print("----------------------------------------------------\n")
	fmt.Printf("Get your tickets here to attend the most awaited %v. \n", conferenceName)
	fmt.Print("----------------------------------------------------\n")

	var userFirstName string
	var userLastName string
	var userTickets uint
	var email string

	fmt.Println("Please enter your first name.")
	fmt.Scan(&userFirstName)
	fmt.Print("----------------------------------------------------\n")
	fmt.Println("Please enter your last name.")
	fmt.Scan(&userLastName)
	fmt.Print("----------------------------------------------------\n")
	fmt.Println("Please enter the number of ticket you want to purchase.")
	fmt.Scan(&userTickets)
	fmt.Print("----------------------------------------------------\n")
	fmt.Println("Please enter your email address.")
	fmt.Scan(&email)
	fmt.Print("----------------------------------------------------\n")

	fmt.Printf("Thankyou %v %v for booking %v tickets. You will soon receive an email at %v for the confirmation of the booking.\n", userFirstName, userLastName, userTickets, email)
	fmt.Print("----------------------------------------------------\n")

}
