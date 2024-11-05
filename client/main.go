package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if (err!=nil) {
		fmt.Println("Error during the tcp handshake", err)
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		mess, err := reader.ReadString('\n')
		if (err!=nil) {
			fmt.Println("Error reading the task from the server")
			return
		}

		trimed := strings.TrimSpace(mess)
		parsed, err := strconv.Atoi(trimed)
		if (err!=nil) {
			fmt.Println("Error converting string to int", "\"" + trimed +"\"")
			continue
		}
		result := parsed * parsed

		time.Sleep(5*time.Second)
		
		fmt.Fprintf(conn, "square of %d is %d\n", parsed, result)
	}
}