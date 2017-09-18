// +build linux, 386, darwin

package main

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

var host string
var start_port int
var end_port int
var wg sync.WaitGroup

func main() {
	fmt.Println("Host> ")
	fmt.Scan(&host)
	fmt.Println("Starting Port (i.e. 80)> ")
	fmt.Scan(&start_port)
	fmt.Println("End Port (i.e. 8080)> ")
	fmt.Scan(&end_port)
	fmt.Println("Running scan... ")

	runtime.GOMAXPROCS(runtime.NumCPU())

	numberOfPorts := (end_port - start_port) + 1
	if numberOfPorts < 1 {
		fmt.Println("No ports to scan")
		os.Exit(1)
	}
	portsToScan := make(chan int, numberOfPorts)

	go func() {
		for i := start_port; i <= end_port; i++ {
			portsToScan <- i
		}
		close(portsToScan)
	}()

	scanners := runtime.NumCPU() * 8
	if scanners > numberOfPorts {
		scanners = numberOfPorts
	}

	fmt.Printf("Running %d scanners on %d cores with %d ports\n", scanners, runtime.NumCPU(), numberOfPorts)
	for i := 0; i < scanners; i++ {
		wg.Add(1)
		go checkPort(i, host, portsToScan)
	}
	wg.Wait()
}

func checkPort(id int, host string, ports <-chan int) {
	for port := range ports {
		qualified_host := fmt.Sprintf("%s%s%d", host, ":", port)
		// give the connection time to setup. Local port scan works quick but over the internet takes longer
		conn, err := net.DialTimeout("tcp", qualified_host, 1*time.Second) // Got the timeout code from: https://stackoverflow.com/questions/37294052/golang-why-net-dialtimeout-get-timeout-half-of-the-time
		if err != nil {
			if strings.HasSuffix(err.Error(), "too many open files") {
				fmt.Printf("You started too many workers! Unable to open a connection. Increase the number of open files on your OS\n")
				continue
			}
			if strings.HasSuffix(err.Error(), "connection refused") {
				//fmt.Printf("Worker: %d, Port: %d%s\n", id, port, "closed")
				continue
			}
			if strings.HasSuffix(err.Error(), "i/o timeout") {
				//fmt.Printf("Worker: %d, Port: %d%s\n", id, port, "timeout")
				continue
			}
			fmt.Println(err)
			continue
		}
		fmt.Fprintf(conn, "GET / HTTP/1.1\r\nHost: %s\r\n\r\n", host)
		// data travel takes time so we give it 300ms to arrive
		conn.SetReadDeadline(time.Now().Add(300 * time.Millisecond))

		buff := make([]byte, 60)
		n, _ := conn.Read(buff)
		fmt.Printf("Worker: #%d, Port: %d Output: %s\n", id, port, strings.Replace(string(buff[:n]), "\r\n", " ", -1))
	}
	wg.Done()
}
