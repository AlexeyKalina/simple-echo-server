package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s [tcp|udp] host:port\n", os.Args[0])
		os.Exit(1)
	}
	prot := os.Args[1]
	serv := os.Args[2]              // берем адрес сервера из аргументов командной строки
	conn, _ := net.Dial(prot, serv) // открываем TCP-соединение к серверу
	go copyTo(os.Stdout, conn)      // читаем из сокета в stdout
	copyTo(conn, os.Stdin)          // пишем в сокет из stdin
}

func copyTo(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
