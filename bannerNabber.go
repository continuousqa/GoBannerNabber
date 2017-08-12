package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

var host string
var start_port int
var end_port int

func main() {
	user_input()
}

func check_port(host string, start_port, end_port int) {

	for i := start_port; i <= end_port; i++ {
		fmt.Println(i)
		qualified_host := fmt.Sprintf("%s%s%d", host, ":", i)
		conn, err := net.DialTimeout("tcp", qualified_host, 50*time.Millisecond)  // Got the timeout code from: https://stackoverflow.com/questions/37294052/golang-why-net-dialtimeout-get-timeout-half-of-the-time
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n") 
		status, err := bufio.NewReader(conn).ReadString('\n')
		fmt.Println(status)
	}
}

func user_input() {
	fmt.Println("Host> ")
	fmt.Scan(&host)
	fmt.Println("Starting Port (i.e. 80)> ")
	fmt.Scan(&start_port)
	fmt.Println("End Port (i.e. 8080)> ")
	fmt.Scan(&end_port)
	fmt.Println("Running scan... ")
	check_port(host, start_port, end_port)
}
