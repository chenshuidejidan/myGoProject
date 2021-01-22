package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9999")
	defer conn.Close()
	if err != nil {
		fmt.Errorf("dail error")
		return
	}
	var s string
	for {
		go func() {
			io.Copy(os.Stdout, conn)
		}()

		fmt.Scan(&s)
		_, err := conn.Write([]byte(s))
		if err != nil {
			fmt.Errorf("server error")
			return
		}
	}
}
