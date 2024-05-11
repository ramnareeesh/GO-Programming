// my implementation

package main

import (
	"fmt"
	"sync"
	"time"
)

type RequestMessage struct {
	ID             int
	SequenceNumber int
}

type TokenMessage struct {
	LN     []int
	Queue  []*Site
	Status string
	Owner  int
}

type Site struct {
	ID     int
	RN     []int
	inbox  chan RequestMessage
	outbox chan TokenMessage
}

var mutex sync.Mutex
var Token TokenMessage
var wg sync.WaitGroup

func contains(lst []*Site, val *Site) bool {
	for _, item := range lst {
		if item.ID == val.ID {
			return true
		}
	}
	return false
}

func broadcast(sites []*Site, requesterID int, msg RequestMessage) {
	for i, s := range sites {
		if i != requesterID {
			s.inbox <- msg
			fmt.Printf("Site %d received a request from site %d with sequence number %d\n", i, msg.ID, msg.SequenceNumber)
		}
	}
}

func (s *Site) RequestCS(sites []*Site) {
	s.RN[s.ID] += 1
	msg := RequestMessage{s.ID, s.RN[s.ID]}
	broadcast(sites, s.ID, msg)
}

func (s *Site) HandleRequest(msg RequestMessage, sites []*Site) {
	defer wg.Done()
	if msg.SequenceNumber > s.RN[msg.ID] {
		s.RN[msg.ID] = msg.SequenceNumber
		fmt.Printf("Site %d updated its RN[%d] to %d\n", s.ID, msg.ID, msg.SequenceNumber)
	}

	if (Token.Owner == s.ID) && (Token.Status == "idle") && (s.RN[msg.ID] == Token.LN[msg.ID]+1) {
		mutex.Lock()

		Token.Owner = msg.ID
		Token.Status = "active"
		fmt.Printf("Site %d: Token changed\n", s.ID)
		sites[msg.ID].outbox <- Token

		fmt.Printf("Site %d passed the token to site %d\n", s.ID, msg.ID)

		mutex.Unlock()
	}
}

func (s *Site) ExecuteCS() {
	if Token.Owner == s.ID && Token.Status == "active" {
		fmt.Printf("Site %d is executing the critical section...\n", s.ID)
		time.Sleep(500 * time.Millisecond)
	}
}

func (s *Site) ExitingCS(sites []*Site) {
	if Token.Owner == s.ID && Token.Status == "active" {
		fmt.Printf("Site %d completed executing the critical section\n", s.ID)
		Token.Status = "idle"
		Token.LN[s.ID] = s.RN[s.ID]

		for _, val := range sites {
			if s.RN[val.ID] == Token.LN[val.ID]+1 {
				if !contains(Token.Queue, val) {
					Token.Queue = append(Token.Queue, val)
					fmt.Printf("Site %d added to queue\n", val.ID)
				}
				fmt.Print("The queue is: ")
				for _, item := range Token.Queue {
					fmt.Printf("%d  ", item.ID)
				}
				fmt.Println()
			}
		}

		if len(Token.Queue) != 0 {
			nextSite := Token.Queue[0]
			// delete top
			Token.Queue = Token.Queue[1:]
			// send token to top element
			Token.Owner = nextSite.ID
			Token.Status = "active"
			nextSite.outbox <- Token
			fmt.Printf("Site %d passed token to %d by exiting CS\n", s.ID, nextSite.ID)

		} else {
			fmt.Printf("No sites requesting for token ... \n")
		}

	}
}

func main() {

	var numSites int
	fmt.Print("Enter the number of sites: ")
	fmt.Scanf("%d", &numSites)

	sites := make([]*Site, numSites)
	for i := range sites {
		sites[i] = &Site{
			ID:     i,
			RN:     make([]int, len(sites)),
			inbox:  make(chan RequestMessage, len(sites)-1),
			outbox: make(chan TokenMessage, len(sites)-1),
		}
	}

	var initiatorSite int
	fmt.Print("Enter the ID of the site that has the token (initially): ")
	fmt.Scanf("%d", &initiatorSite)

	Token.Owner = initiatorSite
	Token.Status = "idle"
	Token.LN = make([]int, len(sites))

	var orderlen int
	fmt.Print("Enter the # of the sites you want to simulate: ")
	fmt.Scanf("%d", &orderlen)

	order := make([]int, orderlen)
	fmt.Print("Enter the order in which sites request the critical section (site IDs separated by space): ")
	for i := range order {
		fmt.Scan(&order[i])
	}

	for _, val := range order {
		site := sites[val]
		go site.RequestCS(sites)

		for _, othersites := range sites {
			if othersites.ID != site.ID {
				msg := <-othersites.inbox
				wg.Add(1)
				go othersites.HandleRequest(msg, sites)

			}

		}
	}
	wg.Wait()

	for _, val := range order {
		site := sites[val]
		select {
		case <-site.outbox:
			{
				fmt.Println("")
				site.ExecuteCS()
				site.ExitingCS(sites)
			}
		default:
			fmt.Printf("Site %d does not have the token\n", site.ID)
		}
	}

	//time.Sleep(3 * time.Second)

}
