package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

type Client struct {
	Conn     net.Conn
	Name     string
	Messages chan string
}

func NewClient(conn net.Conn) *Client {
	return &Client{
		Conn:     conn,
		Messages: make(chan string),
	}
}

func (c *Client) Read(chat *Chat) {
	defer c.Conn.Close()
	for {
		msg, err := bufio.NewReader(c.Conn).ReadString('\n')
		quit := strings.ToLower(msg)
		if err != nil || quit == "quit\n" {
			log.Print(c.Name + " has left our chat...\n")
			chat.Broadcast(c, c.Name+" has left our chat...\n")
			c.Delete(chat)
			break
		}
		msg = strings.Trim(msg, "\r\n")
		if msg == "" {
			continue
		}
		msg = "[" + c.Name + "]:" + msg + "\n"
		currentTime := time.Now()
		msg = "[" + currentTime.Format("2006-01-02 15:04:05") + "]" + msg
		c.Messages <- msg
	}
}

func (client *Client) NameClient(chat *Chat) {
	client.Msg("[ENTER YOUR NAME]: ")

	name, err := bufio.NewReader(client.Conn).ReadString('\n')
	errHandleLogPrint(err, "failed to read name\n")
	name = strings.Trim(name, "\r\n")

	for name == "" || client.isExists(chat, name) {
		if name == "" {
			client.Msg("Please enter at least one symbol\n")
		} else {
			client.Msg("Given name already exists, please try another one\n")
		}
		name, err = bufio.NewReader(client.Conn).ReadString('\n')
		errHandleLogPrint(err, "failed to read name\n")
		name = strings.Trim(name, "\r\n")
	}

	client.Name = name
	// Вывод приветственного сообщения с именем
	client.Msg("Welcome, " + name + "!\nTo exit the chat, please type \"quit\".\n")
}

func (c *Client) Msg(msg string) {
	c.Conn.Write([]byte(msg))
}

func (c *Client) Run(chat *Chat) {
	for msg := range c.Messages {
		chat.Broadcast(c, msg)
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		formattedMsg := fmt.Sprintf("[%s][%s]:", currentTime, c.Name)
		c.Msg(formattedMsg)
	}
}

func (c *Client) ShowHistory(history []string) {
	for _, msg := range history {
		c.Msg(msg)
	}
}

func (c *Client) isExists(chat *Chat, newName string) bool {
	for _, client := range chat.Clients {
		if client.Name == newName {
			return true
		}
	}
	return false
}

func (c *Client) Delete(chat *Chat) {
	for i, client := range chat.Clients {
		if client.Name == c.Name {
			chat.Clients = append(chat.Clients[:i], chat.Clients[i+1:]...)
			break
		}
	}
}

func errHandleLogPrint(err error, msg string) {
	if err != nil {
		log.Print(msg, "error:", err.Error())
	}
}
