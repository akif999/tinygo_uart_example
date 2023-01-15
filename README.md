# UART examples for tinygo

## Summary of UARTs

| item | paramters |
| ---- | ---- |
| wires | 2 |
| speed | 9600, 19200, 38400, 57600, 115200 .. 1500000 |
| spped error tolerance | generally within 10% |
| transmission method | asynchronous |
| max number of transmitter | 1 |
| max number of receiver | 1 |

## Why do high baud rates cause dropouts?

### example case

For example, the problem occurred in the following configuration.

* transmitter
    * Adafruit Feather-m4-express
* receiver
    * Seeed Xiao

test results is this.

| baudrate | result |
| ---- | ---- |
| 9600 | ok |
| 19200 | ok |
| 38400 | ok |
| 57600 | ok |
| 115200 | ng |
| 230400 | ng |

This was a problem in the user's receiver code.

It does not works.
```rx.go
	for {
		if uart.Buffered() > 0 {
			data, _ := uart.ReadByte()
			print(string(data))
		}
		time.Sleep(1 * time.Millisecond)

	}

```

It works.
```rx.go
	for {
		if uart.Buffered() > 0 {
			data, _ := uart.ReadByte()
			print(string(data))
		}
		time.Sleep(500 * time.Microsecond)

	}

```

In other words, it is necessary to appropriately adjust the reception 
polling timing according to the baud rate. 
For example, in the case of 115,200 kbps, a dropout occurred in processing every 500us,  
and no dropout occurred in processing every 100us.


## References

* https://www.analog.com/jp/analog-dialogue/articles/uart-a-hardware-communication-protocol.html
