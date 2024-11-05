package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"strings"
	"time"
)

func dispatchTask(conn net.Conn) {
	defer func (){
		conn.Close()
	}()
	
	reader := bufio.NewReader(conn)
	for {
		randNum := rand.Intn(100)

		fmt.Fprintln(conn, randNum)
		
		responce, err := reader.ReadString('\n')
		if (err!=nil) {
			fmt.Println("Error reading the responce from the client", err)
			return
		}
		responce = strings.TrimSpace(responce)

		fmt.Println(responce)

		time.Sleep(2*time.Second)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if (err!=nil) {
		fmt.Println("Error creating the server", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server running on port 8080...")

	for {
		conn, err := listener.Accept()
		if (err!=nil) {
			fmt.Println("Error during the tcp handshake", err)
			continue
		}
		
		go dispatchTask(conn)
	}
}