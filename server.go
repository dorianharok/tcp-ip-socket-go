package main

import (
	"fmt"
	"log"
	"net"
)

func server() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("클라이언트 문자 수신: ", string(buf[:n]))
	_, err = conn.Write([]byte("Hello, World"))
	if err != nil {
		log.Println(err)
		return
	}
}

func main() {
	server()
}
