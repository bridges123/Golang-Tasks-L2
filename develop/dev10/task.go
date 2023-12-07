package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"
)

/*
=== Утилита telnet ===
Реализовать простейший telnet-клиент.

Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Требования:
Программа должна подключаться к указанному хосту (ip или доменное имя + порт) по протоколу TCP. После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s)
При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться. При подключении к несуществующему сервер, программа должна завершаться через timeout
*/

func main() {
	host := flag.String("host", "", "Host address (ip or domain)")
	port := flag.Int("port", 23, "Port number")
	timeout := flag.Duration("timeout", 10*time.Second, "Connection timeout")
	flag.Parse()

	if *host == "" {
		fmt.Println("Host is required.")
		os.Exit(1)
	}

	address := fmt.Sprintf("%s:%d", *host, *port)
	fmt.Printf("Connecting to %s...\n", address)

	conn, err := net.DialTimeout("tcp", address, *timeout)
	if err != nil {
		fmt.Println("Error connecting to the server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Connected!")

	go func() {
		buffer := make([]byte, 1024)
		for {
			n, err := conn.Read(buffer)
			if err != nil {
				fmt.Println("Connection closed by the server.")
				os.Exit(0)
			}
			fmt.Print(string(buffer[:n]))
		}
	}()

	go func() {
		buffer := make([]byte, 1024)
		for {
			n, err := os.Stdin.Read(buffer)
			if err != nil {
				fmt.Println("Error reading from STDIN:", err)
				os.Exit(1)
			}
			_, err = conn.Write(buffer[:n])
			if err != nil {
				fmt.Println("Error writing to the server:", err)
				os.Exit(1)
			}
		}
	}()

	// Обработка Ctrl+C
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan

	fmt.Println("\nCtrl+C received. Closing the connection.")
}
