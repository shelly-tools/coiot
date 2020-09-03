### CoIoT Library written in Golang (based on CoAP package) for the Shelly IoT-Devices

Package shelly-tools/coiot provides coiot (CoAP) communication with the popular Shelly IoT-Devices from Allterco Robotics.

#### Installation
Install via `$ go get github.com/shelly-tools/coiot`.

Please add -u flag to update in the future.

#### Examples

A basic client example to query `/cit/d` path from the Shelly:
```golang
package main

import (
	"log"
	"os"

	"github.com/shelly-tools/coiot"
)

func main() {
	req := coiot.Message{
		Type:      coiot.Confirmable,
		Code:      coiot.GET,
		MessageID: 12,
		Payload:   []byte("Hello Shelly"),
	}
	req.SetPathString("/cit/d")
	c, err := coiot.Dial("udp", "192.168.178.212:5683")
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
```

please check the example directory for more..

#### Credits

This library is based on the following packages:

* https://github.com/dustin/go-coap
* https://github.com/bulyshko/coiot

Special "Thanks" to Dustin Sallings & Romuald Bulyshko who did most of the work.