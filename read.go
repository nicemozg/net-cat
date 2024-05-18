package main

import (
	"fmt"
	"log"
	"os"
)

func ReadArg() string {
	arg := os.Args[1:]

	if len(arg) > 1 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		os.Exit(0)
	}

	port := "8989"

	if len(arg) == 1 {
		port = arg[0]
		for _, c := range port {

			if c > '9' || c < '0' {
				fmt.Println("[USAGE]: ./TCPChat $port")
				os.Exit(0)
			}
		}
	}
	return port
}

func Greating() string {
	dat, err := os.ReadFile("greating.txt")
	if err != nil {
		log.Println("Failed to load greating")
		return ""
	}
	return string(dat)
}
