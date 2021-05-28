package main

import (
	"fmt"
	"net"
	"strings"
)

func connHandler(c net.Conn) {
	if c == nil {
		return
	}
	buf := make([]byte, 4096)
	for {
		n, err := c.Read(buf)
		if err != nil || n == 0 {
			c.Close()
			break
		}
		inStr := strings.TrimSpace(string(buf[0:n]))
		inputs := strings.Split(inStr, " ")
		switch inputs[0] {
		case "ping":
			c.Write([]byte("pong\n"))
		case "echo":
			echoStr := strings.Join(inputs[1:], " ") + "\n"
			c.Write([]byte(echoStr))
		case "quit":
			c.Close()
			break
		default:
			fmt.Printf("Unsupported command: %s\n", inputs[0])
		}
	}
	fmt.Printf("Connection from %v closed. \n", c.RemoteAddr())
}
func main() {
	server, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Printf("Fail to start server, %s\n", err)
	}
	fmt.Println("Server Started ...")
	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Printf("Fail to connect, %s\n", err)
			break
		}
		go connHandler(conn)
	}
}
