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

func mqttGoroutine() {
	clientId := "tinygo-client-" + randomString(10)
	println("ClientId:", clientId)
	println("Selected MQTT broker at", broker)
	// We'll reuse the client across connections.
	mqttClient = mqtt.NewClient(mqtt.ClientConfig{
		Decoder: mqtt.DecoderNoAlloc{UserBuffer: make([]byte, 1500)},
		OnPub: func(_ mqtt.Header, _ mqtt.VariablesPublish, r io.Reader) error {
			message, _ := io.ReadAll(r)
			println("Message", string(message), "received on topic", topic)
			return nil
		},
	})
	pubVars := mqtt.VariablesPublish{
		TopicName:        []byte(topic),
		PacketIdentifier: uint16(randomInt(0, 65000)),
	}
	var connVars mqtt.VariablesConnect
	connVars.SetDefaultMQTT([]byte(clientId))
	connVars.KeepAlive = 60 // seconds; some brokers reject KeepAlive=0
	pubflags, err := mqtt.NewPublishFlags(mqtt.QoS0, false, false)
	if err != nil {
		panic("failed to create correct pubflags")
	}
	for {
		time.Sleep(2 * time.Second)
		// Retry connecting to broker until success.
		println("Dialing TCP to", broker, "...")
		conn, err := net.Dial("tcp", broker)
		if err != nil {
			println("dialing", broker, "failed, retring in a bit...")
			continue
		}

		println("TCP connected", conn.RemoteAddr().String())
		err = mqttClient.Connect(context.Background(), conn, &connVars)
		if err != nil {
			println("failed to connect to broker:", err.Error())
			conn.Close()
			continue
		}
		for mqttClient.IsConnected() {
			println("publishing MQTT message...")
			pubVars.PacketIdentifier++
			data := "{\"e\":[{ \"dv\":" +
				strconv.Itoa(int(dialValue)) +
				", \"bp\":" +
				strconv.FormatBool(buttonPush) +
				", \"tp\":" +
				strconv.FormatBool(touchPush) +
				" }]}"
			err = mqttClient.PublishPayload(pubflags, pubVars, []byte(data))
			if err != nil {
				println("failed to publish", err.Error())
			}
			time.Sleep(time.Second)
		}
		println("MQTT client disconnected, retrying...")
		if mqttClient.Err() != nil {
			println("disconnect reason:", mqttClient.Err().Error())
		}
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
