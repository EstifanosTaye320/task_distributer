package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"net"
	"sync"
	"time"
)

var lstClients = make(map[net.Conn]bool)
var mu sync.Mutex

func readAndWrite(conn net.Conn) {
	reader := bufio.NewReader(conn)

	for {
		responce, err := reader.ReadString('\n')
		if (err!=nil) {
			fmt.Println("Error reading the responce from the client", err)
			continue
		}

		fmt.Println(responce)
	}
} 

func dispatchTask(conn net.Conn) {
	defer func (){
		mu.Lock()
		delete(lstClients, conn)
		mu.Unlock()
		conn.Close()
	}()

	go readAndWrite(conn)

	for {
		randNum := rand.IntN(100)
		fmt.Fprintln(conn, randNum)

		time.Sleep(time.Duration(2)*time.Second)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if (err!=nil) {
		fmt.Println("Error creating the server", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if (err!=nil) {
			fmt.Println("Error during the tcp handshake", err)
			continue
		}

		mu.Lock()
		lstClients[conn] = true
		mu.Unlock()
		
		go dispatchTask(conn)
	}
}