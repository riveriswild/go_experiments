package main

import (
	"fmt"
	"go_tutorial/helper"
	"time"
	"sync"
)

var conferenceName = "Go Conference"
const conferenceTickets uint = 50
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	userName string
	lastName string
	email string
	userTickets uint
}

var wg = sync.WaitGroup{}

func main () {
	greetUsers(remainingTickets)

	for {

		userName, lastName, email, userTickets := getUserInput()

		// bookings[0] = userName + " " + lastName -- for arrays
		isValidName, isValidEmail, validTicketCount := helper.ValidateUserInput(userName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && validTicketCount {
			remainingTickets := bookTicket(remainingTickets, userTickets, userName, lastName, email)

			wg.Add(1)
            go sendTicket(userTickets, userName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("These are the bookings: %v\n", firstNames)


			var noTicketsRemaining bool = remainingTickets == 0

			if noTicketsRemaining {
				fmt.Println("Our conference is booked out")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Invalid name")
			} 
			if !isValidEmail {
				fmt.Println("No @ in email")
			}
			if !validTicketCount {
				fmt.Println("Invalid ticket num")
			}
		}
	}
	wg.Wait()
}

func greetUsers(remainingTickets uint) {
	fmt.Printf("Welcome to our conference %v!\n", conferenceName)
	fmt.Printf("Welcome to %v booking app \n", conferenceName)
	fmt.Printf("Please book tickets! There were %v tickets \n", conferenceTickets)
	fmt.Println(remainingTickets, "are still available")
}

func getFirstNames() [] string {

	firstNames := []string{}
	for _, booking := range bookings {
		// var names = strings.Fields(booking)
		// var firstName = names[0]
		firstNames = append(firstNames, booking.userName)
	fmt.Printf("These are the bookings: %v\n", firstNames)
	}
	return firstNames
}


func getUserInput() (string, string, string, uint){
	var userName string
	var userTickets uint
	var lastName string
	var email string


	// ask for name
	fmt.Println("Enter your name")
	fmt.Scan(&userName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets")
	fmt.Scan(&userTickets)

	return userName, lastName, email, userTickets
}

func bookTicket(remainingTickets uint, userTickets uint, userName string, lastName string, email string) (uint){
	var userData = UserData {
		userName: userName,
		lastName: lastName,
		email: email,
		userTickets: userTickets,
	}
	bookings = append(bookings, userData)


	remainingTickets = remainingTickets - userTickets
	fmt.Printf("User %v %v with email %v booked %v tickets \n", userName, lastName, email, userTickets)
	fmt.Printf("Remaining tickets: %v \n", remainingTickets)
	return remainingTickets
}

func sendTicket(userTickets uint, userName string, lastName string, email string) {
	time.Sleep(10*time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, userName, lastName)
	fmt.Println("******************")
	fmt.Printf("Sending tickets %v to %v", ticket, email)
	fmt.Println("******************")
	wg.Done()

}