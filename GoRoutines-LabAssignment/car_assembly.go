package main

import (
	"fmt"
	"time"
)

func engineSetup(done chan bool) {
	fmt.Println("Engine setup started")
	time.Sleep(3 * time.Second)
	fmt.Println("Engine setup complete")
	done <- true
}

func acceleratorBrakeClutchSetup(done chan bool) {
	fmt.Println("Accelerator, brake, and clutch setup started")
	time.Sleep(3 * time.Second)
	fmt.Println("Accelerator, brake, and clutch setup complete")
	done <- true
}

func bodyAssembling(done chan bool) {
	fmt.Println("Body assembling started")
	time.Sleep(3 * time.Second)
	fmt.Println("Body assembling complete")
	done <- true
}

func steeringSetup(done chan bool) {
	fmt.Println("Steering setup started")
	time.Sleep(3 * time.Second)
	fmt.Println("Steering setup complete")
	done <- true
}

func main() {
	engineDone := make(chan bool)
	acceleratorBrakeClutchDone := make(chan bool)
	bodyDone := make(chan bool)
	steeringDone := make(chan bool)

	go engineSetup(engineDone)
	<-engineDone

	go acceleratorBrakeClutchSetup(acceleratorBrakeClutchDone)
	<-acceleratorBrakeClutchDone

	go bodyAssembling(bodyDone)
	<-bodyDone

	go steeringSetup(steeringDone)
	<-steeringDone

	fmt.Println("Car assembly complete")
}
