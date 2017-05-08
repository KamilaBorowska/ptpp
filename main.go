package main

import (
	"bufio"
	"encoding/base64"
	"log"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {
	reader := bufio.NewReader(conn)
	message, _ := reader.ReadString('\n')
	parts := strings.Split(message, " ")
	contents := "File content"
	if strings.TrimSpace(parts[2]) == "PTTPU/1.0" {
		contents = base64.StdEncoding.EncodeToString([]byte(contents))
	}
	conn.Write([]byte(contents + "\n<<PTTP END>>\n"))
	conn.Close()
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}
