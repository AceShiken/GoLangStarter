package helper

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var threadNo = 1

func Greetings(conferenceName string, conferenceTickets int, remainingTickets int) {
	fmt.Println("Welcome to", conferenceName, "booking app")
	fmt.Printf("We have %v out of %v available tickets.\n", remainingTickets, conferenceTickets)
}

// function or variables with small casing are package access modifier
func generateTicket(tickets int, userName string) string {
	var ticket = fmt.Sprintf("%v TICKETS for %v", tickets, userName)
	return ticket
}

// function or variables with large casing are public access modifier
func SendTicket(tickets int, userName string) {
	wg.Add(threadNo)
	var generatedTicket = generateTicket(tickets, userName)
	fmt.Println("\nSending ticket", generatedTicket, "to", userName)
	time.Sleep(10 * time.Second)
	fmt.Println("\nTickets Sent to", userName)
	wg.Done()
}

func WaitSending() {
	wg.Wait()
}
