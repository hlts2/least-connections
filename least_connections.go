package leastconnections

import (
	"net/url"
	"sync"

	"github.com/pkg/errors"
)

// ErrServersNotExist is the error that servers dose not exists
var ErrServersNotExist = errors.New("servers dose not exist")

// LeastConnections --
type LeastConnections interface {
	Next() (next *url.URL, done func())
}

type leastConnections struct {
	urls  []*url.URL
	conns map[int]int
	mu    *sync.Mutex
}

// New initializes a new instance of LeastConnected
func New(urls []*url.URL) (LeastConnections, error) {
	if len(urls) == 0 {
		return nil, ErrServersNotExist
	}

	conns := make(map[int]int)
	for i := range urls {
		conns[i] = 0
	}

	return &leastConnections{
		urls:  urls,
		conns: conns,
		mu:    new(sync.Mutex),
	}, nil
}

func (lc *leastConnections) Next() (*url.URL, func()) {
	var (
		min = -1
		idx int
	)

	lc.mu.Lock()

	for urlIdx, cnt := range lc.conns {
		if min == -1 || cnt < min {
			min = cnt
			idx = urlIdx
		}
	}

	lc.conns[idx]++

	lc.mu.Unlock()

	return lc.urls[idx], func() {
		lc.mu.Lock()
		lc.conns[idx]--
		lc.mu.Unlock()
	}
}
