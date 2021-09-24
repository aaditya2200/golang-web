package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panicln(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		io.WriteString(conn, "\nHello From Our Server\n")
		fmt.Fprintln(conn, "Whats up?")
		fmt.Fprintf(conn, "%v", "Yes sir")

		conn.Close()
	}
}
