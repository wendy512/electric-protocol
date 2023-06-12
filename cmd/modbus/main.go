package main

import (
	"fmt"
	"github.com/goburrow/serial"
	modbus "github.com/thinkgos/gomodbus/v2"
	"time"
)

func main() {
	p := modbus.NewRTUClientProvider(modbus.WithEnableLogger(),
		modbus.WithSerialConfig(serial.Config{
			Address:  "COM3",
			BaudRate: 9600,
			DataBits: 8,
			StopBits: 1,
			Parity:   "N",
			Timeout:  modbus.SerialDefaultTimeout,
		}))

	client := modbus.NewClient(p)
	if err := client.Connect(); err != nil {
		panic(err)
	}

	defer client.Close()

	fmt.Println("starting")
	for {
		// send: 01 04 00 01 00 01 60 0a 湿度值：01 04 02 01 86 39 02
		// sen: 01 04 00 00 00 01 31 ca

		//温度
		// _, err := client.ReadInputRegistersBytes(1, 0, 1)

		//湿度
		//_, err := client.ReadInputRegistersBytes(1, 1, 1)

		//温度+湿度
		_, err := client.ReadInputRegistersBytes(1, 0, 2)
		if err != nil {
			panic(err)
		}

		//fmt.Printf("ReadDiscreteInputs %#v\r\n", results)
		time.Sleep(time.Second * 2)
	}
}
