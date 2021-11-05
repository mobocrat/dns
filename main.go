package main

import (
	"fmt"
	"net"
	"time"
)

const address = "8.8.8.8:53"

func main() {
	sock, err := net.Dial("udp", address)
	if err != nil {
		panic(err)
	}
	defer sock.Close()
	// TODO: Use command-line args to replace 'google.com' in the following packet
	req := []byte{
		0xdb, 0x42, // ID
		1, 0, // Flags
		0, 1, // QDCOUNT
		0, 0, // ANCOUNT
		0, 0, // NSCOUNT
		0, 0, // ARCOUNT
		6,
		'g', 'o', 'o', 'g', 'l', 'e',
		3,
		'c', 'o', 'm',
		0,
		0, 1, // A Record
		0, 1,
	}
	if _, err := sock.Write(req); err != nil {
		panic(err)
	}
	resp := make([]byte, 1024)
	sock.SetReadDeadline(time.Now().Add(5 * time.Second))
	n, err := sock.Read(resp)
	if err != nil {
		panic(err)
	}
	// TODO: Parse response data to extract IP
	for i := 0; i < n; i += 8 {
		fmt.Printf("%0x", resp[i:i+8])
		fmt.Print(" ")
		fmt.Printf("%s", string(resp[i:i+8]))
		fmt.Print("\n")
	}
}
