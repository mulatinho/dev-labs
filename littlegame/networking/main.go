package networking

import (
	"fmt"
	"net"
	"time"

	"github.com/mulatinho/golabs/littlegame/utils"
)

const (
	NETWORKING_SINGLE_PLAYER = iota
	NETWORKING_MULTI_PLAYER
)

type Client struct {
	Connection net.Conn
	Address    string
}

type Server struct {
	ListenAddress string
	Listener      net.Listener
	Clients       []net.Conn
}

func NewServer(addr string) *Server {
	return &Server{
		ListenAddress: addr,
	}
}

func (s *Server) StartLoop() error {
	var err error

	if s.Listener, err = net.Listen("tcp", s.ListenAddress); err != nil {
		return err
	}

	go s.serverLoop()

	return nil
}

func (s *Server) serverLoop() {
	for {
		time.Sleep(time.Millisecond * 200)
		newConnection, err := s.Listener.Accept()
		if err != nil {
			fmt.Printf("error Accept(): %+v\n", err)
			continue
		}

		go s.clientLoop(newConnection)
	}
}

func (s *Server) clientLoop(connection net.Conn) {
	utils.LogMessage(utils.LOG_TYPE_DEBUG, connection)
	s.Clients = append(s.Clients, connection)

	select {}
}
