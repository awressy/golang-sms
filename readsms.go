package main

import (
	"fmt"
	"github.com/tarm/serial"
	"time"
	"log"
	// "strings"
)

func main(){
	detect()
}

func detect(){
	c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 115200, ReadTimeout: 1}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 1024)
	// s.Write([]byte("AT+CMGR=11\r\n"))
	// time.Sleep(1000 * time.Millisecond)
	// // data, _ := s.Write([]byte("AT+CMGL=\"all\"\r"))
	// data, _ = s.Read(buf)
	// data := string(buf[:data])
	// time.Sleep(1000 * time.Millisecond)
	// // data, _ := s.Read(buf)

	// fmt.Println(data)
	// // s.Write([]byte(message + "\u001A"))

	s.Write([]byte("AT+CMGF=1\r"))

	time.Sleep(5000 * time.Millisecond)

	n1, _ := s.Write([]byte("AT+CMGL=\"all\"\r\n"))

	n1, _ = s.Read(buf)

	sos1 := string(buf[:n1])

	fmt.Println(sos1)
}