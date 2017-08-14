package main

import (
	"fmt"
	"net"
	"time"
	"sync"
)

var host string
var start_port int
var end_port int
var wg sync.WaitGroup

func main() {
	user_input()
}

func check_port(host string, start_port, end_port int) {

	for i := start_port; i <= end_port; i++ {
		//fmt.Println('\n')
		qualified_host := fmt.Sprintf("%s%s%d", host, ":", i)
		conn, err := net.DialTimeout("tcp", qualified_host, 10*time.Millisecond)  // Got the timeout code from: https://stackoverflow.com/questions/37294052/golang-why-net-dialtimeout-get-timeout-half-of-the-time
		if err != nil {
			continue
		}
		fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n1\n22\n\n\n\n")
		conn.SetReadDeadline(time.Now().Add(10*time.Millisecond))

		// swapping bufio reads to reading buffers as bytes gets huge performance increase: 8000 ports in 20s (vs 1min 10s using bufio reads)
		buff := make([]byte, 1024)
		n, _ := conn.Read(buff)
		fmt.Printf("Port: %d%s\n",i, buff[:n])
	}
	wg.Done()
}

func user_input() {
	fmt.Println("Host> ")
	fmt.Scan(&host)
	fmt.Println("Starting Port (i.e. 80)> ")
	fmt.Scan(&start_port)
	fmt.Println("End Port (i.e. 8080)> ")
	fmt.Scan(&end_port)
	fmt.Println("Running scan... ")

	//check_port_set1(host, start_port, end_port) // 15s to run 1000 ports sequentially

	port_range := end_port - start_port
	end_port_set1 := (port_range / 4) + start_port
	end_port_set2 := (port_range / 4) + end_port_set1
	end_port_set3 := (port_range / 4) + end_port_set2

	wg.Add(4)		// 3s to run 1000 ports on 4 concurrent groups
	go check_port(host, start_port, end_port_set1)
	go check_port(host, (end_port_set1 + 1), end_port_set2)
	go check_port(host, (end_port_set2 + 1), end_port_set3)
	go check_port(host, (end_port_set3 + 1), end_port)
	wg.Wait()

}
