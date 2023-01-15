// This reads from UART1 and outputs to default serial, usually UART0 or USB.
// Example of how to work with UARTs other than the default.
package main

import (
	"machine"
	"time"
)

var (
	uart = machine.UART1
	tx   = machine.UART_TX_PIN
	rx   = machine.UART_RX_PIN
)

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	go func() {
		for {
			led.Low()
			time.Sleep(time.Millisecond * 1000)
			led.High()
			time.Sleep(time.Millisecond * 1000)
		}
	}()

	uart.Configure(machine.UARTConfig{BaudRate: 115200, TX: tx, RX: rx})
	for {
		if uart.Buffered() > 0 {
			data, _ := uart.ReadByte()
			print(string(data))
		}
		time.Sleep(100 * time.Microsecond)

	}
}
