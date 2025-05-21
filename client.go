package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func client() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	fmt.Print("Send Message: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	conn.Write([]byte(input))
	var buf = make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("서버 응답:", string(buf[:n]))
}

func main() {
	client()
}
