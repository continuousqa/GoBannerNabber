package main

import (
	"bufio"
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

func check_port_set1(host string, start_port, end_port int) {

	for i := start_port; i <= end_port; i++ {
		//fmt.Println(i)
		qualified_host := fmt.Sprintf("%s%s%d", host, ":", i)
		conn, err := net.DialTimeout("tcp", qualified_host, 50*time.Millisecond)  // Got the timeout code from: https://stackoverflow.com/questions/37294052/golang-why-net-dialtimeout-get-timeout-half-of-the-time
		if err != nil {
			//fmt.Println("set1: ", err)
			continue
		}
		fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n1\n22\n\n\n\n")
		conn.SetReadDeadline(time.Now().Add(50*time.Millisecond))

		status, err := bufio.NewReader(conn).ReadString('\n')
		fmt.Println(i, status)
	}
	wg.Done()
}

func check_port_set2(host string, start_port, end_port int) {

	for i := start_port; i <= end_port; i++ {
		//fmt.Println(i)
		qualified_host := fmt.Sprintf("%s%s%d", host, ":", i)
		conn, err := net.DialTimeout("tcp", qualified_host, 50*time.Millisecond)  // Got the timeout code from: https://stackoverflow.com/questions/37294052/golang-why-net-dialtimeout-get-timeout-half-of-the-time
		if err != nil {
			//fmt.Println("set2: ", err)
			continue
		}
		fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n1\n22\n\n\n\n")
		conn.SetReadDeadline(time.Now().Add(50*time.Millisecond))

		status, err := bufio.NewReader(conn).ReadString('\n')
		fmt.Println(i, status)
	}
	wg.Done()
}
func check_port_set3(host string, start_port, end_port int) {

	for i := start_port; i <= end_port; i++ {
		//fmt.Println(i)
		qualified_host := fmt.Sprintf("%s%s%d", host, ":", i)
		conn, err := net.DialTimeout("tcp", qualified_host, 50*time.Millisecond)  // Got the timeout code from: https://stackoverflow.com/questions/37294052/golang-why-net-dialtimeout-get-timeout-half-of-the-time
		if err != nil {
			//fmt.Println("set3: ", err)
			continue
		}
		fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n1\n22\n\n\n\n")
		conn.SetReadDeadline(time.Now().Add(50*time.Millisecond))

		status, err := bufio.NewReader(conn).ReadString('\n')
		fmt.Println(i, status)
	}
	wg.Done()
}
func check_port_set4(host string, start_port, end_port int) {

	for i := start_port; i <= end_port; i++ {
		//fmt.Println(i)
		qualified_host := fmt.Sprintf("%s%s%d", host, ":", i)
		conn, err := net.DialTimeout("tcp", qualified_host, 50*time.Millisecond)  // Got the timeout code from: https://stackoverflow.com/questions/37294052/golang-why-net-dialtimeout-get-timeout-half-of-the-time
		if err != nil {
			//fmt.Println("set4: ", err)
			continue
		}
		fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n1\n22\n\n\n\n")
		conn.SetReadDeadline(time.Now().Add(50*time.Millisecond))

		status, err := bufio.NewReader(conn).ReadString('\n')
		fmt.Println(i, status)
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
	//end_port_set1 := (port_range / 2) + start_port
	end_port_set1 := (port_range / 4) + start_port
	end_port_set2 := (port_range / 4) + end_port_set1
	end_port_set3 := (port_range / 4) + end_port_set2

	//wg.Add(2)		// 13s to run 1000 ports on 2 concurrent groups
	//go check_port_set1(host, start_port, end_port_set1)
	//go check_port_set2(host, (end_port_set1 + 1), end_port)
	//wg.Wait()

	wg.Add(4)		// 9s to run 1000 ports on 4 concurrent groups
	go check_port_set1(host, start_port, end_port_set1)
	go check_port_set2(host, (end_port_set1 + 1), end_port_set2)
	go check_port_set3(host, (end_port_set2 + 1), end_port_set3)
	go check_port_set4(host, (end_port_set3 + 1), end_port)
	wg.Wait()

}
