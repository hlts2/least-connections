package leastconnections

import (
	"log"

	"github.com/pkg/errors"

	"github.com/hlts2/lock-free"
)

// ErrServersNotExist is the error that servers dose not exists
var ErrServersNotExist = errors.New("servers dose not exist")

// ErrServerNotExist is the error that server dose not exists
var ErrServerNotExist = errors.New("server dose not exist")

// Servers is custom type of servers
type Servers []string

type LeastConnections interface {
	IncrementConnections(server string)
	DecrementConnections(server string)
	Next() string
}

type leastConnections struct {
	servers     Servers
	connections map[string]int
	lf          lockfree.LockFree
}

// New initializes a new instance of LeastConnected
func New(servers Servers) (LeastConnections, error) {
	if len(servers) == 0 {
		return nil, ErrServersNotExist
	}

	connections := make(map[string]int)
	for _, server := range servers {
		connections[server] = 0
	}

	return &leastConnections{
		servers:     servers,
		connections: connections,
		lf:          lockfree.New(),
	}, nil
}

// IncrementConnection increments the number of connection of server
func (lc *leastConnections) IncrementConnections(server string) {
	lc.lf.Wait()

	if _, ok := lc.connections[server]; ok {
		lc.connections[server]++
	}

	lc.lf.Signal()
}

// DecrementConnection decrements the number of connection of server
func (lc *leastConnections) DecrementConnections(server string) {
	lc.lf.Wait()

	if v, ok := lc.connections[server]; ok {
		if v > 0 {
			lc.connections[server]--
		}
	}

	lc.lf.Signal()
}

func (lc *leastConnections) Next() string {
	lc.lf.Wait()

	var minConnectionsServer string
	minConnections := -1

	var (
		cnt int
		ok  bool
	)
	for _, server := range lc.servers {
		if cnt, ok = lc.connections[server]; !ok {
			log.Fatal(errors.WithMessage(ErrServerNotExist, server))
		}

		if minConnections == -1 || cnt < minConnections {
			minConnections = cnt
			minConnectionsServer = server
		}
	}

	lc.lf.Signal()

	return minConnectionsServer
}
