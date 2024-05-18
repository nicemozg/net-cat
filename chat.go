package main

import "time"

const (
	maxConn = 10
)

type Chat struct {
	MaxConn int
	Clients []*Client
	History []string
}

func NewChat() *Chat {
	chat := &Chat{
		MaxConn: maxConn,
		Clients: make([]*Client, 0),
		History: make([]string, 0),
	}
	return chat
}

func (c *Chat) Broadcast(sender *Client, msg string) {
	joinMessage := sender.Name + " has joined our chat...\n"
	if msg != joinMessage {
		c.History = append(c.History, msg)
	}

	for _, m := range c.Clients {
		currentTime := time.Now()

		newMessage := "\n" + msg + "[" + currentTime.Format("2006-01-02 15:04:05") + "][" + m.Name + "]:"
		if sender.Conn.RemoteAddr() != m.Conn.RemoteAddr() {
			m.Msg(newMessage)
		}
	}
}
