# Xiao ESP32-C3 Sensor Station

## What you need

    - Seeed Studio Xiao ESP32-C3 board
    - Grove IoT Starter Kit parts
    - Personal computer with Go 1.25+ and TinyGo installed, and a serial port.

### TinyGo drivers

All of the code dependencies you will need are already in the Go modules file in this directory, so they will be downloaded and installed automatically. You don't need to do anything, when you follow the subsequent instructions they will be downloaded by TinyGo.

Just for your information, the TinyGo drivers that let you connect to sensors, displays, and other external peripheral devices are located in the separate repository at https://github.com/tinygo-org/drivers


## Connecting the Seeed Studio Xiao ESP32-C3 to your computer

<img src="./assets/xiao-esp32c3-pins.webp" alt="Seeed Studio Xiao ESP32-C3" width="500"/>

Plug the Seeed Studio Xiao ESP32-C3 into your computer using a USB cable. There may be one provided in your starter kit.

You should also attach the small antenna to the Xiao ESP32-C3 board uFL connector for best reception.

## Running the code

The TinyGo programs will run directly on the Xiao microcontoller. The procedure is basically:

- Edit your TinyGo program.
- Compile and flash it to your Xiao.
- The program executes from the Xiao. You can disconnect the Xiao from your computer and plug it into a battery if you wish, the program executes directly on the microcontroller.

Let's get started!

## Code

### step1.go - Green LED

![Seeed Studio Xiao ESP32-C3](./assets/step1.jpg)

This tests that you can compile and flash your Xiao with TinyGo code by blinking an LED.

We will use a separate LED that we will connect from the parts kit.

- Connect one of the "Ground" pins on the Xiao to the breadboard's ground rail (-) using a black or brown jumper cable.

- Connect the "3.3V" pin on the Xiao to the breadboard's power rail (+) using a red jumper cable.

- Plug the Grove green LED into the provided cable with the Grove connector on one end, and the male jumpers on the other. Make sure the LED itself is plugged into the Grove board.

- Connect the black male end of the Grove cable to the breadboard's ground rail (-).

- Connect the red male end of the Grove cable to the breadboard's power rail (+).

- Connect the yellow male end of the Grove cable to pin D1 on the Xiao.


Run the code.

```
tinygo flash -target xiao-esp32c3 -size short ./step1/
```

You should see the green LED blink.


### step2.go - Green LED, Button

![Xiao](./assets/step2.jpg)

Next, we will attach a push button. The button will control the LED we added in the previous step. When you push the button, the LED will turn on. When you release the button, it will turn off.

- Plug the Grove Button into a provided cable with the Grove connector on one end, and the male jumpers on the other.

- Connect the black male end of the Grove cable to the breadboard's ground rail (-).

- Connect the red male end of the Grove cable to the breadboard's power rail (+).

- Connect the yellow male end of the Grove cable to pin D10 on the Xiao.


Run the code.

```
tinygo flash -target xiao-esp32c3 ./step2/
```

When you press the button, the green LED should turn on.


### step3.go - Green LED, Button, Red LED

![Xiao](./assets/step3.jpg)

Now we will add a second LED, that will will also control with the same push button. One LED will be on with the button is pushed, and the other will be on when the button is released.

- Plug the Grove red LED into one of the provided cable with the Grove connector on one end, and the male jumpers on the other.

- Connect the black male end of the Grove cable to the breadboard's top left set of pins (-).

- Connect the red male end of the Grove cable to the breadboard's top right (+) set of pins.

- Connect the yellow male end of the Grove cable to pin D8 on the Xiao.

Run the code.

```
tinygo flash -target xiao-esp32c3 ./step3/
```

The red LED should light up. When you press the button, the green LED should turn on, and the red LED should turn off. When you release the button, the green LED should turn off, and the red LED should turn on again.

### step4.go - Green LED, Button, Red LED, Buzzer, Touch

![Xiao](./assets/step4.jpg)

In this step we will add two new devices. The first one is a capacitive touch sensor. It essentially acts like a button, but you only need to touch it to activate it. The second is a small piezoelectric buzzer. When current is sent to the buzzer, it makes a noise.

- Plug the Grove buzzer into one of the provided cables with the Grove connector on one end, and the male jumpers on the other.

- Connect the black male end of the Grove cable to the breadboard's top left set of pins (-).

- Connect the red male end of the Grove cable to the breadboard's top right (+) set of pins.

- Connect the yellow male end of the Grove cable to pin D2 on the Xiao.

- Plug the Grove touch sensor into one of the provided cables with the Grove connector on one end, and the male jumpers on the other.

- Connect the black male end of the Grove cable to the breadboard's top left set of pins (-).

- Connect the red male end of the Grove cable to the breadboard's top right (+) set of pins.

- Connect the yellow male end of the Grove cable to pin D3 on the Xiao.

Run the code.

```
tinygo flash -target xiao-esp32c3 ./step4/
```

When you touch the touch sensor, the buzzer should emit a noise.


### step5.go - Green LED, Button, Red LED, Buzzer, Touch, Dial

![Xiao](./assets/step5.jpg)

The next device we will add is a rotary dial. The rotary dial is an analog device. The amount of voltage that is allowed to pass through it is based on how far you turn the dial. 

In order to read whatever voltage is going thru the dial, we will need to use an [Analog to Digital Converter (ADC)](https://en.wikipedia.org/wiki/Analog-to-digital_converter). The Xiao Nano RP 2040 connect has several ADC onboard for us to use.

We will use this to control the brightness of the red LED. To do this, we will need to switch from just a digital on or off signal, to use [Pulse Width Modulation (PWM)](https://en.wikipedia.org/wiki/Pulse-width_modulation). We can use PWM to make the LED brighter or dimmer.

- Plug the Grove dial (Rotary Angle Sensor) into one of the provided cables with the Grove connector on one end, and the male jumpers on the other.

- Connect the black male end of the Grove cable to the breadboard's top left set of pins (-).

- Connect the red male end of the Grove cable to the breadboard's top right (+) set of pins.

- Connect the yellow male end of the Grove cable to pin D0 on the Xiao.


Run the code.

```
tinygo flash -target xiao-esp32c3 ./step5/
```

Adjusting the dial sensor should control the brightness of the red LED.


### step6.go - Green LED, Button, Red LED, Buzzer, Touch, Dial, OLED display

![Xiao](./assets/step6.jpg)

Now we will add a SSD1306 OLED display to show the state of the buttons and dial. We will control this display using an [I2C interface](https://en.wikipedia.org/wiki/I%C2%B2C).

- Connect a jumper wire from the "GND" pin on the breadboard next to the OLED display, to the breadboard's top left set of pins (-).

- Connect a jumper wire from the "VCC" pin on the breadboard next to the OLED display, to the breadboard's top right (+) set of pins.

- Connect a jumper wire from the "SDA" pin on the breadboard next to the OLED display, to the Xiao D4 pin.

- Connect a jumper wire from the "SCL" pin on the breadboard next to the OLED display, to the Xiao D5 pin.


We have 2 TinyGo packages to make it easier to use small displays such as the SSD1306 in the kit. 

The TinyFont package renders fonts to any of the supported displays in the TinyGo drivers repo. 

The TinyDraw package has a number of drawing primitives like circles, lines, and triangles that can be used with any of the supported displays.

Run the code.

```
tinygo flash -target xiao-esp32c3 ./step6/
```

The dial should now cause the OLED display to show its current position. The OLED should also have two empty circles that will light up when you press the Button to light up the Blue LED and when you press the touch sensor respectively.

### step7.go - Green LED, Button, Red LED, Buzzer, Touch, Dial, OLED display - Alarm Logic

![Xiao](./assets/step6.jpg)

Now that we have assembled all of the hardware for the sensor device, let's implement some logic to make it into an alarm system. 

We now want it to work like this: 

- when the button is pushed, the alarm system is made active if it is off, and turned off if it is on. the green LED should light up to indicate if the alarm system is active.
- when the rotary dial is turned above a certain level, the alarm buzzer should sound. This indicates that the "power level has been exceeded".
- when the touch sensor is touched, the red LED should light up, and the alarm buzzer should sound. this is like a test of the system.

Take look at the changes to the code to implement this in step 7. Now try flashing the updated code and trying it out.

```
tinygo flash -target xiao-esp32c3 ./step7/
```

### step8.go - Alarm System Web Server - LED control

![Xiao Web Server](./assets/xiao-http-webserver.png)

In this step we will connect to the built-in wireless radio on the ESP32-C3 and set it up to act as both a WiFi Access Point and a web server.

The code for the wireless connection is in the file `wifi.go`.

We will run a tiny web server on the board itself, that is intended to allow the user to manage the device using a web browser. No additional hardware is required for this step.

Take a look at the code in the file `webserver.go`. A few things of note:

- The HTML page is embedded into the binary that is loaded onto the Xiao using the normal `go:embed` directive.
- Likewise the CSS file is also embedded into the binary. The CSS is from a very small framework called [MinCSS](https://mincss.com/docs.html).
- Even though it is running on the Xiao, the normal `net/html` package is used to implement the web server using calls to `http.HandleFunc()` and `http.ListenAndServe()`.

The Xiao board will act as an access point. You will need to connect your computer to it to connect to the web server running on the Xiao. Replace `myssid` with something unique for your WiFi setup, then flash the board with the following command:

```
tinygo flash -target xiao-esp32c3 -ldflags="-X main.ssid=myssid" -size short -monitor ./step8/
```

Your should see something in your output similar to this:

```
Device reset.
Connected to /dev/ttyACM0. Press Ctrl-C to exit.
SHA-256 comparison failed:
Calculated: d6cadb12af152d93b79bfab7ea444967a8ffb62d0214e60b38afb0c19355b0d8
Expected: 669e977affaaeb08afc2331378d727825998a37dd98c06ccd2bb44865694a4af
Attempting to boot anyway...
entry 0x4039f48c
Connecting to WiFi...
HTTP server listening on http://192.168.1.1:80
```

#### How to tell if it is working

![Webserver 1](./assets/alarm-webserver-1.png)

As long as your computer is connected to the access point running on the Xiao, you can open a new browser window and connect to the IP address shown in your terminal from the Xiao board's output. For example, http://192.168.1.1:80

Click on the "On" button on the web page to activate the alarm system. Click on the "Off" button to deactivate it.

### step9.go - Alarm System Web Server - display alarm status

![Xiao](./assets/step6.jpg)

Now we will modify the web server so we can see the alarm system status on the web page.

Remember to replace `myssid` for your WiFi setup, then flash the board with the following command:

```
tinygo flash -target xiao-esp32c3 -ldflags="-X main.ssid=myssid" -size short -monitor ./step9/
```

#### How to tell if it is working

![Webserver 2](./assets/alarm-webserver-2.png)

Connect your computer to the access point, then open a new browser window and connect to the IP address shown in your terminal. For example,  http://192.168.1.1:80

### step10.go - Alarm System Web Server - Set alarm limit

![Xiao](./assets/step6.jpg)

Now we will modify the web server so we can set the limit for the rotary dial to trigger the alarm from the web page.

Remember to replace `myssid` for your WiFi setup, then flash the board with the following command:

```
tinygo flash -target xiao-esp32c3 -ldflags="-X main.ssid=myssid" -size short -monitor ./step10/
```

#### How to tell if it is working

![Webserver 3](./assets/alarm-webserver-3.png)

Connect your computer to the access point, then open a new browser window and connect to the IP address shown in your terminal. For example, http://192.168.4.1:8080/

### mqtt.go (bonus step) - Green LED, Button, Red LED, Buzzer, Touch, Dial, OLED display, MQTT

![Xiao](./assets/step6.jpg)

In this step we will connect to a machine to machine messaging server using the [MQTT machine to machine messaging protocol](https://en.wikipedia.org/wiki/MQTT). No additional hardware is required for this step.

However, now instead of the Xiao acting as a web server, instead it will act as an MQTT client.

You will need access to the Internet, so you must connect the Xiao board as a WiFi client to some other WiFi Access Point.

Substitute the correct values for your WiFi setup in the following command:

```
tinygo flash -target xiao-esp32c3 -ldflags="-X main.ssid=myssid -X main.password=mypass" -size short -monitor ./mqtt
```

## How to tell if it is working

Install the client tools for `mosquitto` for your operating system, then run:

```
mosquitto_sub -h 'broker.hivemq.com' -t 'tinygohackday'
```

When you run this command, you should be able to see the messages appear. These are being sent from your own machine to an MQTT broker provided by the Eclipse Foundation and running on a cloud server that they provide.
