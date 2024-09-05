package main

// go mod init tcp_proxy
// go mod tidy
// go build -o tcp_proxy
// ./tcp_proxy

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8082") // Changed to 8082
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	log.Println("TCP proxy listening on 127.0.0.1:8082")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(clientConn net.Conn) {
	defer clientConn.Close()

	backendConn, err := net.Dial("tcp", "192.168.10.30:80")
	if err != nil {
		log.Println(err)
		return
	}
	defer backendConn.Close()

	go io.Copy(backendConn, clientConn)

	scanner := bufio.NewScanner(backendConn)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "IP: 192.168.10.30") {
			line = strings.Replace(line, "IP: 192.168.10.30", "IP: 185.28.22.166", 1)
		}
		if strings.HasPrefix(line, "RemoteAddr: 192.168.10.1:") {
			line = strings.Replace(line, "RemoteAddr: 192.168.10.1:", "RemoteAddr: 185.28.22.166:", 1)
		}
		fmt.Fprintf(clientConn, "%s\r\n", line)
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
