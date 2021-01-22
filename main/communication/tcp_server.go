package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:9999")
	if err != nil {
		fmt.Errorf("net listen error")
		return
	}
	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Errorf("listen error")
			return
		}

		go handleConnect(conn)
	}
}

func handleConnect(conn net.Conn) {
	buf := make([]byte, 10)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Errorf("read error")
			break
		}
		fmt.Println(string(buf[0:n]))
		fmt.Println("read total", n, " character")
		conn.Write([]byte(strings.ToUpper(string(buf[0:n]))))
	}
	conn.Close()
}
