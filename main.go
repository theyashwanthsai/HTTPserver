package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)


func handleIncomingRequest(conn net.Conn){
	defer conn.Close()

	reader := bufio.NewReader(conn)

	requestline,err := reader.ReadString('\n')
	if err != nil{
		fmt.Println("Error readin req:", err)
		return
	}

	parts := strings.Fields(requestline)
	if len(parts) != 3{
		fmt.Println("Invalid request format")
		return
	}

	method := parts[0]
	path := parts[1]

	res := "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\n\r\nHello, this is a simple Go web server!"

	conn.Write([]byte(res))

	fmt.Printf("Req - [%s] %s\n", method, path)
}


func main(){
	port := 8900
	
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d"))
	if err != nil{
		fmt.Println("Error starting server:", err)
		return
	}

	fmt.Printf("Server listening on port %d...\n", port)

	for {
		conn, err := listener.Accept()
		if err != nil{
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleIncomingRequest(conn)
	}
}