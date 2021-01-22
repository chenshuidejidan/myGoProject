package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		n, err := io.WriteString(conn, time.Now().Format("2006-01-02 15:04:05\n"))
		if err != nil {
			log.Print(conn.RemoteAddr(), " 出现问题，正在退出....")
			return
		}
		log.Print("成功向", conn.RemoteAddr(), "写入", n, "个字节数据")
		time.Sleep(3 * time.Second)
	}
}
