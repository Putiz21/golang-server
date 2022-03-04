package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	host = "127.0.0.1"
	port = "2137"
	TCP  = "tcp"
)

func main() {
	//Socket
	fmt.Println("Trying to bind TCP Server on: " + host + ":" + port)
	listener, err := net.Listen(TCP, host+":"+port)
	if err != nil {
		fmt.Println("Failed to listen on: ", host+":"+port+" Error: "+err.Error())
	}
	defer listener.Close()

	//Połączenie
	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println("Failed to connect: ", err.Error())
		}
		fmt.Println("Client Connected from: ", connection.LocalAddr().String())
		//Konsola
		for {
			go odbior(connection)
			shell := bufio.NewReader(os.Stdin)
			fmt.Print("# ")
			input, errr := shell.ReadString('\n')
			if errr != nil {
				fmt.Println("Shell Error: " + errr.Error())
			}
			buffer := []byte(input)
			connection.Write(buffer)
		}
	}
}

//odbior
func odbior(connection net.Conn) {
	buffer, err := bufio.NewReader(connection).ReadBytes('\n')
	if err != nil {
		fmt.Println("Client disconnected from: " + connection.LocalAddr().String())
		connection.Close()
		os.Exit(1)
	}
	fmt.Println("\n" + string(buffer[:len(buffer)-1]))
	fmt.Print("# ")
}
