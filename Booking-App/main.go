package main

import (
	"booking-app/helper"
	"fmt"
)

var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets = 50

type UserData struct {
	userName        string
	numberOfTickets int
}

func main() {

	// %T -> variable type printing
	// fmt.Printf("conferenceName : %T, conferenceTickets : %T, reminaingTickets : %T\n", conferenceName, conferenceTickets, remainingTickets)

	helper.Greetings(conferenceName, conferenceTickets, remainingTickets)

	var bookings = make(map[string]int)
	var userData = make([]UserData, 0)

	for remainingTickets > 0 {
		userName, userTickets := getUserInput()

		fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

		if userTickets > remainingTickets {
			fmt.Printf("Only %d tickets available\n", remainingTickets)
			continue
		}

		bookings, remainingTickets, userData = bookTicket(bookings, userName, userTickets, userData)
		fmt.Printf("Number of tickets remaining %d\n", remainingTickets)
		// Async threads process with go routines
		go helper.SendTicket(userTickets, userName)
	}
	fmt.Println("SOLDOUT!!")
	// _ -> is blank identifier
	for userName, tickets := range bookings {
		fmt.Printf("User %v booked: %d\n", userName, tickets)
	}

	for _, user := range userData {
		fmt.Println(user)
	}
	// waits to close main thread umtil all child threads are completed
	helper.WaitSending()

}

func bookTicket(bookings map[string]int, userName string, userTickets int, userData []UserData) (map[string]int, int, []UserData) {
	bookings[userName] = userTickets
	userData = append(userData, UserData{userName: userName, numberOfTickets: userTickets})
	remainingTickets -= userTickets
	return bookings, remainingTickets, userData
}

func getUserInput() (string, int) {
	var userName string
	var userTickets int

	// Scan into address
	fmt.Println("Enter username")
	fmt.Scan(&userName)

	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets)
	return userName, userTickets
}

// References : https://www.youtube.com/watch?v=yyUHQIec83I
// build -> 'go mod init booking-app'
// run -> 'go run main.go helper.go' or 'go run .'
