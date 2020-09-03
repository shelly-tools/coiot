package main

import (
	"log"
	"os"

	"github.com/shelly-tools/coiot"
)

func main() {

	//Send CoIoT Exec to Shelly RGBW2 with "Listen CoAp for color change commands" enabled.
	ip := "192.168.178.70"
	path := "/cit/e"
	payload := "{\"a\":100,\"i\":[101,102,103,104,105,106,107,108],\"v\":[254,0,0,0,100,100,100,1]}"
	if len(os.Args) > 1 {
		ip = os.Args[1]
		path = os.Args[2]
		payload = os.Args[3]
	}

	req := coiot.Message{
		Type:      coiot.Confirmable,
		Code:      coiot.EXEC,
		MessageID: 12345,
		Payload:   []byte(payload),
	}
	req.SetPathString(path)

	c, err := coiot.Dial("udp", ip + ":5683")
	if err != nil {
		log.Fatalf("Error dialing: %v", err)
	}

	rv, err := c.Send(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}

	if rv != nil {
		log.Println("Payload successfully sent!")
	}
}
