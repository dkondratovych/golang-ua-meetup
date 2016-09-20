package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	// TCP listener
	go func() {
		ln, err := net.Listen("tcp", ":8081")
		if err != nil {
			log.Print(err.Error())
		}

		conn, err := ln.Accept()

		for {
			message, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				log.Print(err.Error())
			}

			log.Printf("Message: %s", message)

			_, err := conn.Write([]byte("Bazinga!"))
			if err != nil {

			}
			return
		}
	}()

	time.Sleep(2 * time.Second)

	// TCP Dialer
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(2*time.Second))
		defer cancel()

		dialer := &net.Dialer{}
		conn, _ := dialer.DialContext(ctx, "tcp", ":8081")
		//conn, _ := net.Dial("tcp", ":8081").Dia
		for {
			fmt.Fprintf(conn, "Hi there \n")
		}

		//dialer := &net.Dialer{}
		//conn, err := dialer.Dial("tcp", "127.0.0.1:8081")
		//if err != nil {
		//	log.Print(err.Error())
		//}

	}()

	time.Sleep(2 * time.Second)
}
