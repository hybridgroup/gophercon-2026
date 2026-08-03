package main

import (
	"time"

	nl "tinygo.org/x/drivers/netlink"
	"tinygo.org/x/espradio/netlink"
)

var (
	ssid     string
	password string
	port     string = ":80"

	link netlink.Esplink
)

func connectWifi() {
	// wait a bit for serial
	time.Sleep(2 * time.Second)

	println("Connecting to WiFi...")
	err := link.NetConnect(&nl.ConnectParams{
		Ssid:       ssid,
		Passphrase: password,
	})

	if err != nil {
		failMessage("could not connect to WiFi: " + err.Error())
	} else {
		println("Connected to Wifi!")
	}
}

func failMessage(msg string) {
	for {
		println(msg)
		time.Sleep(1 * time.Second)
	}
}
