package main

import "time"

func main() {
	go initializeRfid()
	
	startRfid()
	// 如何过五秒之后停止rfid
	time.Sleep(5 * time.Second)
	stopRfid()
}
