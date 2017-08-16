// +build linux, 386, darwin

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

	port_range := end_port - start_port
	end_port_set1 := (port_range / 10) + start_port
	end_port_set2 := (port_range / 10) + end_port_set1
	end_port_set3 := (port_range / 10) + end_port_set2
	end_port_set4 := (port_range / 10) + end_port_set3
	end_port_set5 := (port_range / 10) + end_port_set4
	end_port_set6 := (port_range / 10) + end_port_set5
	end_port_set7 := (port_range / 10) + end_port_set6
	end_port_set8 := (port_range / 10) + end_port_set7
	end_port_set9 := (port_range / 10) + end_port_set8


	wg.Add(10)		// 1min to run 65,000 ports on 10 concurrent groups
	go check_port(host, start_port, end_port_set1)
	go check_port(host, (end_port_set1 + 1), end_port_set2)
	go check_port(host, (end_port_set2 + 1), end_port_set3)
	go check_port(host, (end_port_set3 + 1), end_port_set4)
	go check_port(host, (end_port_set4 + 1), end_port_set5)
	go check_port(host, (end_port_set5 + 1), end_port_set6)
	go check_port(host, (end_port_set6 + 1), end_port_set7)
	go check_port(host, (end_port_set7 + 1), end_port_set8)
	go check_port(host, (end_port_set8 + 1), end_port_set9)
	go check_port(host, (end_port_set9 + 1), end_port)
	wg.Wait()

}
