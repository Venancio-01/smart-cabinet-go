package main

import (
	"fmt"
	"log"

	"go.bug.st/serial"
)

var port serial.Port

func initializeRfid() {
	mode := &serial.Mode{
		BaudRate: 115200,
	}

	var err error
	port, err = serial.Open("/dev/ttyS3", mode)
	if err != nil {
		log.Fatal(err)
	}
}

func openRfidListener() {
	buff := make([]byte, 100)
	for {
		n, err := port.Read(buff)
		if err != nil {
			log.Fatal(err)
			break
		}
		if n == 0 {
			fmt.Println("\nEOF")
			break
		}
		fmt.Printf("%v", string(buff[:n]))

		messageQueue.Push(string(buff[:n]))
	}
}

func startRfid() {
	const COMMAND_START = "5A"

	commandBody := "000102100008" + generateAntennaCommand() + "01020006"
	checkCode := generateCRC16Code(commandBody)
	command := []byte(COMMAND_START + commandBody + checkCode)

	port.Write(command)

	fmt.Printf("Start Rfid\n")
}

func stopRfid() {
	port.Write([]byte{0x5A, 0x00, 0x01, 0x01, 0x01, 0x00, 0x00, 0x00, 0xEB, 0xD5})

	fmt.Printf("Stop Rfid\n")
}
