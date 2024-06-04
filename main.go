package main

import "time"

func main() {
	initializeRfid()
	startRfid()
	go openRfidListener()

	time.Sleep(5 * time.Second)
	stopRfid()
}
