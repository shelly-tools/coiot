package main

import (
	"flag"
	"log"
	"time"

	"github.com/shelly-tools/coiot"
)

var (
	Id uint16 = 12345
)

func main() {

	//Send CoIoT Exec to Shelly RGBW2 with "Listen CoAp for color change commands" enabled.
	//payload := ""
	up := true
	IpAddress := flag.String("ip", "224.0.1.187", "the Shellys ip address")
	Interval := flag.Int("interval", 2000, "Interval to change the color")
	loopcount := flag.Int("loops", 10, "Number of loops")
	brightness := flag.Int("brightness", 100, "brightness")
	gain := flag.Int("gain", 100, "gain")
	Transition := flag.String("transition", "1000", "Interval to change the color")

	Pcol := flag.String("pcol", "0,0,255,0", "primary color rbgw, default is 0,0,255,0")
	Scol := flag.String("scol", "255,0,0,0", "secondary color in rgbw, default is 255,0,0,0")

	flag.Parse()
	// Simple infinite color change between blue and green every 5 seconds
	for i := 0; i < *loopcount; i++ {
		var payload string

		if up == true {
			up = false
			payload = `{"a":100,"i":[101,102,103,104,105,106,107,108],"v":[` + *Pcol + `,50,10,` + *Transition + `,1]}`
		} else {
			up = true
			payload = `{"a":100,"i":[101,102,103,104,105,106,107,108],"v":[` + *Scol + `,10,10,` + *Transition + `,1]}`
		}
		Id = Id + 1
		req := coiot.Message{
			Type:      coiot.NonConfirmable,
			Code:      coiot.EXEC,
			MessageID: Id,
			Payload:   []byte(payload),
		}
		req.SetPathString("/cit/e")

		c, err := coiot.Dial("udp", *IpAddress+":5683")
		if err != nil {
			log.Fatalf("Error dialing: %v", err)
		}
		rv, _ := c.Send(req)

		if rv != nil {
			log.Println(payload)
		}
		time.Sleep(time.Duration(*Interval) * time.Millisecond)
	}
}
