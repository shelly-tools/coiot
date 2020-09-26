package main

import (
	"flag"
	"log"
	"time"

	"github.com/shelly-tools/coiot"
)

func main() {

	//Send CoIoT Exec to Shelly RGBW2 with "Listen CoAp for color change commands" enabled.
	payload := ""
	ip := flag.String("ip", "224.0.1.187", "the Shellys ip address")
	up := true

	// Simple infinite color change between blue and green every 5 seconds
	for {
		if up == true {
			up = false
			payload = "{\"a\":100,\"i\":[101,102,103,104,105,106,107,108],\"v\":[0,0,255,0,100,100,2000,1]}"
		} else {
			up = true
			payload = "{\"a\":100,\"i\":[101,102,103,104,105,106,107,108],\"v\":[0,255,0,0,100,100,2000,1]}"
		}

		req := coiot.Message{
			Type:      coiot.Confirmable,
			Code:      coiot.EXEC,
			MessageID: 12345,
			Payload:   []byte(payload),
		}
		req.SetPathString("/cit/e")

		c, err := coiot.Dial("udp", *ip+":5683")
		if err != nil {
			log.Fatalf("Error dialing: %v", err)
		}

		rv, _ := c.Send(req)

		if rv != nil {
			log.Println(payload)
		}
		time.Sleep(2000 * time.Millisecond)
	}

}
