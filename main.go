package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	port := ReadArg()
	chat := NewChat()
	greating := Greating()

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}

	defer listener.Close()
	log.Printf("Listening on the port :" + port + "\n")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("failed to accept connection: %s", err.Error())
			continue
		}
		client := NewClient(conn)
		if len(chat.Clients) >= chat.MaxConn {
			client.Msg("Sorry, but chat is too full, try later")
			conn.Close()
			continue
		}

		go func() {
			client.Msg(greating)
			client.NameClient(chat)
			client.ShowHistory(chat.History)
			currentTime := time.Now().Format("2006-01-02 15:04:05")
			formattedMsg := fmt.Sprintf("[%s][%s]:", currentTime, client.Name)
			client.Msg(formattedMsg)
			chat.Broadcast(client, client.Name+" has joined our chat...\n")
			log.Printf("new client has joined: %s", client.Name)

			chat.Clients = append(chat.Clients, client)
			go client.Read(chat)
			go client.Run(chat)
		}()

	}
}
