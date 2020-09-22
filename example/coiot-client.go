package main

import (
	"log"
        "flag"
	"github.com/shelly-tools/coiot"
)

func main() {

	ip := flag.String("ip", "192.168.178.240", "the Shellys ip address")
	path := flag.String("path", "/cit/s", "the CoIoT path - /cit/d or /cit/s")
	payload := flag.String("payload","", "Payload to send")

	flag.Parse()
	req := coiot.Message{
		Type:      coiot.Confirmable,
		Code:      coiot.GET,
		MessageID: 12345,
		Payload:   []byte(*payload),
	}
	req.SetPathString(*path)

	c, err := coiot.Dial("udp", *ip+":5683")
	if err != nil {
		log.Fatalf("Error dialing: %v", err)
	}

	rv, err := c.Send(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}

	if rv != nil {
		log.Printf("Response payload: %s", rv.Payload)
	}

}
