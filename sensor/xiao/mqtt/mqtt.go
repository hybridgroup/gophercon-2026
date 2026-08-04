package main

import (
	"context"
	"io"
	"math/rand"
	"net"
	"strconv"
	"time"

	mqtt "github.com/soypat/natiu-mqtt"
)

func connectToMQTT() {
	clientId := "tinygo-client-" + randomString(10)
	println("ClientId:", clientId)

	// Get a transport for MQTT packets.
	// Retry TCP connection since public brokers may reject/close connections under load.
	println("Connecting to MQTT broker at", broker)
	var conn net.Conn
	for attempt := range 5 {
		var err error
		conn, err = net.Dial("tcp", broker)
		if err != nil {
			println("net.Dial attempt", attempt+1, "failed:", err)
			time.Sleep(2 * time.Second)
			continue
		}
		break
	}
	if conn == nil {
		failMessage("all TCP connection attempts failed")
	}
	println("TCP connected to", conn.RemoteAddr())
	defer conn.Close()

	// Create new client
	mqttClient = mqtt.NewClient(mqtt.ClientConfig{
		Decoder: mqtt.DecoderNoAlloc{UserBuffer: make([]byte, 1500)},
		OnPub: func(_ mqtt.Header, _ mqtt.VariablesPublish, r io.Reader) error {
			message, _ := io.ReadAll(r)
			println("Message", string(message), "received on topic", topic)
			return nil
		},
	})

	// Connect client
	var varconn mqtt.VariablesConnect
	varconn.SetDefaultMQTT([]byte(clientId))
	varconn.KeepAlive = 60 // seconds; some brokers reject KeepAlive=0
	println("Sending MQTT CONNECT...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := mqttClient.Connect(ctx, conn, &varconn)
	if err != nil {
		failMessage("failed to connect: " + err.Error())
	}
	println("MQTT CONNECT succeeded")
}

func publishToMQTT() {
	pubFlags, _ := mqtt.NewPublishFlags(mqtt.QoS0, false, false)
	pubVar := mqtt.VariablesPublish{
		TopicName: []byte(topic),
	}

	for {
		println("Publishing MQTT message...")
		data := "{\"e\":[{ \"dv\":" +
			strconv.Itoa(int(dialValue)) +
			", \"bp\":" +
			strconv.FormatBool(buttonPush) +
			", \"tp\":" +
			strconv.FormatBool(touchPush) +
			" }]}"

		pubVar.PacketIdentifier++
		err := mqttClient.PublishPayload(pubFlags, pubVar, []byte(data))
		if err != nil {
			failMessage("error transmitting message: " + err.Error())
		}
		time.Sleep(time.Second)
	}
}

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// Generate a random string of A-Z chars with len = l
func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(65, 90))
	}
	return string(bytes)
}
