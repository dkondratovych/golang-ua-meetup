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
		if err != nil {
			log.Print(err.Error())
		}

		for {
			message, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				log.Print(err.Error())
			}

			log.Printf("Message: %s", message)
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
	}()

	time.Sleep(2 * time.Second)
}
