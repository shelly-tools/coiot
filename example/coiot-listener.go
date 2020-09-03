package main

import (
	"log"
	"net"
	"os"
	"github.com/shelly-tools/coiot"
)

func CoIoTHandler(l *net.UDPConn, a *net.UDPAddr, m *coiot.Message) *coiot.Message {
	s := string(m.Payload)
	if len(os.Args) > 1 {
		ip := os.Args[1]
		if a.String() == ip +":5683" {
			log.Printf("%s - DeviceType: %s - DeviceID: %s - Path: %s - Payload: %s", a, m.DeviceType(), m.DeviceID(), m.Path(), s)
		}
	}  else {
		log.Printf("%s - DeviceType: %s - DeviceID: %s - Path: %s - Payload: %s", a, m.DeviceType(), m.DeviceID(), m.Path(), s)
	}
	return nil
}

func main() {
	mux := coiot.NewServeMux()
	mux.Handle("/cit/s", coap.FuncHandler(CoIoTHandler))
	log.Fatal(coap.ListenAndServe("udp", "224.0.1.187:5683", mux))
}
