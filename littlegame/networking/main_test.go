package networking

import (
	"net"
	"testing"

	"github.com/mulatinho/golabs/mlt"
)

const (
	listenAddr  = ":6000"
	netProtocol = "tcp"
)

func TestNewServer(t *testing.T) {
	newServer := NewServer(listenAddr)
	mlt.Equals(t, newServer.ListenAddress, listenAddr)

	mlt.Equals(t, newServer.StartLoop(), nil)
}

func TestTCPConnection(t *testing.T) {
	initialMsg := []byte{'H', 'I', 'S', 'E', 'R', 'V'}
	newServer := NewServer(listenAddr)
	newServer.StartLoop()

	conn, err := net.Dial(netProtocol, listenAddr)
	mlt.Equals(t, err, nil)

	bytes, err := conn.Write(initialMsg)
	mlt.Assert(t, bytes > 0)
	mlt.Equals(t, err, nil)
	mlt.Equals(t, len(newServer.Clients), 1)

	select {}
}
