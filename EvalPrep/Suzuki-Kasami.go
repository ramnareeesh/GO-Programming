package main

import (
	"fmt"
	"time"
)

type Message struct {
	Type     string
	Sender   int
	Sequence int
}

//type Token struct {
//	Owner  int
//	LN     []int
//	TokenQ []int
//}
//
//type Site struct {
//	ID       int
//	RN       []int
//	HasToken bool
//	Token    Token
//	Mutex    sync.Mutex
//}

func (s *Site) requestCriticalSection(sites []*Site) {
	s.Mutex.Lock()
	s.RN[s.ID]++
	for _, site := range sites {
		if site.ID != s.ID {
			go s.sendRequest(site)
		}
	}
	s.Mutex.Unlock()
}

func (s *Site) sendRequest(receiver *Site) {
	receiver.Mutex.Lock()
	receiver.RN[s.ID] = max(receiver.RN[s.ID], s.RN[s.ID])
	if receiver.HasToken && receiver.RN[s.ID] == receiver.Token.LN[s.ID]+1 {
		receiver.HasToken = false
		s.HasToken = true
		s.Token = receiver.Token
	}
	receiver.Mutex.Unlock()
}

func (s *Site) executeCriticalSection() {
	fmt.Printf("Site %d is executing the critical section.\n", s.ID)
}

func (s *Site) releaseCriticalSection(sites []*Site) {
	s.Mutex.Lock()
	s.Token.LN[s.ID] = s.RN[s.ID]
	for i, site := range sites {
		if site.ID != s.ID && s.RN[i] == s.Token.LN[i]+1 {
			s.Token.TokenQ = append(s.Token.TokenQ, i)
		}
	}
	if len(s.Token.TokenQ) > 0 {
		nextSiteID := s.Token.TokenQ[0]
		s.Token.TokenQ = s.Token.TokenQ[1:]
		s.HasToken = false
		sites[nextSiteID].HasToken = true
		sites[nextSiteID].Token = s.Token
	}
	s.Mutex.Unlock()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var numSites int
	fmt.Print("Enter the number of sites: ")
	fmt.Scan(&numSites)

	sites := make([]*Site, numSites)
	for i := range sites {
		sites[i] = &Site{
			ID:       i,
			RN:       make([]int, len(sites)),
			HasToken: false,
			Token: Token{
				Owner:  -1,
				LN:     make([]int, len(sites)),
				TokenQ: make([]int, 0),
			},
		}
	}

	var firstSite int
	fmt.Print("Enter the ID of the site that starts with the token: ")
	fmt.Scan(&firstSite)
	sites[firstSite].HasToken = true
	sites[firstSite].Token.Owner = firstSite

	order := make([]int, numSites)
	fmt.Println("Enter the order in which sites request the critical section (site IDs separated by space):")
	for i := range order {
		fmt.Scan(&order[i])
	}

	for _, siteID := range order {
		go func(siteID int) {
			sites[siteID].requestCriticalSection(sites)
			time.Sleep(1 * time.Second)
			sites[siteID].executeCriticalSection()
			sites[siteID].releaseCriticalSection(sites)
		}(siteID)
	}

	time.Sleep(time.Duration(numSites) * time.Second)
}
