package main

import(
	"log"
	//"fmt"
	"github.com/tarm/serial"
	"time"
)

func main(){
	phone := "089687101195"
	message := "Hello dunia"
	port := "/dev/ttyUSB0"
	sendMessage(port, phone, message)
}

func sendMessage(port, phone, message string){
	c := &serial.Config{Name: port, Baud: 115200, ReadTimeout: 1}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	
	s.Write([]byte("AT+CMGF=1\r"))
	
	time.Sleep(100 * time.Millisecond)
	s.Write([]byte("AT+CMGS=\""+phone+"\"\r"))
	time.Sleep(100 * time.Millisecond)
	s.Write([]byte(message + "\u001A"))
}