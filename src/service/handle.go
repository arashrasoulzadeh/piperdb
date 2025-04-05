// Package server /*
// Copyright © 2025 Arash Rasoulzadeh <arashrasoulzadeh@gmail.com>
package service

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/arashrasoulzadeh/piperdb/src/models/commands"
	"github.com/arashrasoulzadeh/piperdb/src/models/server"
	"net"
)

func Accept(s *server.Server) {
	for {
		// Accept incoming connections
		conn, err := s.Listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		fmt.Println("New connection from", conn.RemoteAddr())

		// Handle client connection in a goroutine
		go HandleClient(conn)
	}
}

func HandleClient(conn net.Conn) {
	defer conn.Close()
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	for {
		req, err := rw.ReadString('\n')
		if err != nil {
			rw.WriteString("failed to read input")
			rw.Flush()
			fmt.Println("Error:", err)
			return
		}

		var ReceivedMessage commands.Command

		err = json.Unmarshal([]byte(req), &ReceivedMessage)

		ReceivedMessage.SetClientId(conn.RemoteAddr().String())
		if err != nil {
			fmt.Println("Error Unmarshal:", err)
			rw.WriteString(fmt.Sprintf("Error unmarshal: %s", err.Error()))
		} else {
			RouteCommand(ReceivedMessage)
		}

		rw.Flush()

	}

}
