package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	done := make(chan struct{})
	for {
		go func() {
			io.Copy(os.Stdout, conn)
			log.Println("done")
			done <- struct{}{}
		}()
		mustCopy(conn, os.Stdin)
	}

	<-done //同步操作，等待后台goroutine完成后，main goroutine才退出
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
