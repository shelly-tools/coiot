package main

import (
	"log"
	"net"

	"github.com/shelly-tools/coiot"
)

func CoIoTHandler(l *net.UDPConn, a *net.UDPAddr, m *coiot.Message) *coiot.Message {
	s := string(m.Payload)
	// non-coiot message? just skip..
	if m.OptionDevice() == nil {
		return nil
	}
	log.Printf("%s - DeviceType: %s - DeviceID: %s - Path: %s - Payload: %s", a, m.DeviceType(), m.DeviceID(), m.Path(), s)
	return nil
}

func main() {
	mux := coiot.NewServeMux()
	mux.Handle("/cit/s", coiot.FuncHandler(CoIoTHandler))
	log.Fatal(coiot.ListenAndServeMulticast("udp", "224.0.1.187:5683", mux))
}
