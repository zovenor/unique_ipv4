package main

import (
	"fmt"
	"time"

	"github.com/zovenor/unique_ipv4"
)

func main() {
	ipList := make([]string, 0, unique_ipv4.MaxUint32)
	l := 1000
	for range l {
		ipList = append(ipList, "0.0.0.0")
	}
	for range l {
		ipList = append(ipList, "255.255.255.255")
	}
	fmt.Println("Start")
	startTime := time.Now()
	amount := unique_ipv4.GetTotalUniqueIPv4Addresses(ipList)
	duration := time.Since(startTime)
	fmt.Printf("Time taken: %s; Total: %d; From: %d\n", duration, amount, l)
}
