package main

import (
	"log"
	"os"

	"github.com/shelly-tools/coiot"
)

func main() {

	ip := "192.168.178.212"
	path := "/cit/d"
	payload := ""
	if len(os.Args) > 1 {
		ip = os.Args[1]
		path = os.Args[2]
		payload = os.Args[3]
	}
	req := coiot.Message{
		Type:      coap.Confirmable,
		Code:      coap.GET,
		MessageID: 12345,
		Payload:   []byte(payload),
	}
	req.SetPathString(path)

	c, err := coap.Dial("udp", ip + ":5683")
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
