package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex

type Ticket struct {
	TicketID    string
	EventName   string
	AvailSeats  int
	BookedSeats int
	Price       int
}

func reserve(ticket *Ticket, no_seats int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Println("Reserving tickets...")
	ticket.AvailSeats -= no_seats
	ticket.BookedSeats += no_seats
	fmt.Println("Event name:", ticket.EventName)
	fmt.Println("Available seats now:", ticket.AvailSeats)
	fmt.Println("Booked seats:", ticket.BookedSeats)
	fmt.Println("Reservation Complete.")
	fmt.Println()
	mutex.Unlock()
	wg.Done()
}

func cancel(ticket *Ticket, no_seats int, wg *sync.WaitGroup) {
	mutex.Lock()
	if ticket.BookedSeats < no_seats {
		fmt.Println("Error: Not enough booked seats to cancel.")
		mutex.Unlock()
		wg.Done()
		return
	}
	fmt.Println("Cancelling tickets...")
	ticket.AvailSeats += no_seats
	ticket.BookedSeats -= no_seats
	fmt.Println("Event name:", ticket.EventName)
	fmt.Println("Available seats now:", ticket.AvailSeats)
	fmt.Println("Booked seats:", ticket.BookedSeats)
	fmt.Println("Cancellation Complete.")
	fmt.Println()
	mutex.Unlock()
	wg.Done()
}

func main() {
	Concert := &Ticket{TicketID: "T1", EventName: "Concert", AvailSeats: 100, BookedSeats: 0, Price: 50}
	Movie := &Ticket{TicketID: "T2", EventName: "Movie", AvailSeats: 50, BookedSeats: 0, Price: 30}

	var wg sync.WaitGroup

	wg.Add(4)
	go reserve(Concert, 5, &wg) //3
	go reserve(Movie, 3, &wg)   //1
	go cancel(Movie, 2, &wg)    // 4
	go cancel(Concert, 2, &wg)  //2

	wg.Wait()
	fmt.Println("All transactions complete")
}
