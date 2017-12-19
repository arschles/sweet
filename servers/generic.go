package servers

import (
	"net/http"
	"net/http/httptest"
	"sync"
)

// Generic holds a generic web server that you can use to run your tests against
type Generic interface {
	// Server is the running server. Don't forget to close this!
	Server() *httptest.Server
	// Requests gets the requests that have been sent to the server.
	// The server adds requests to this slice every time it gets one,
	// so if you call this twice you might get more requests in the slice
	// the second time
	Requests() []*http.Request
}

type genericServer struct {
	srv  *httptest.Server
	reqs []*http.Request
	lck  *sync.RWMutex
}

// NewGeneric returns a Generic server (with or without TLS) that responds to all requests
func NewGeneric(tls bool) Generic {
	lck := new(sync.RWMutex)
	reqs := []*http.Request{}
	hdl := http.NewServeMux()
	hdl.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lck.Lock()
		defer lck.Unlock()
		reqs = append(reqs, r)
		w.WriteHeader(http.StatusOK)
	})
	var srv *httptest.Server
	if tls {
		srv = httptest.NewTLSServer(hdl)
	} else {
		srv = httptest.NewServer(hdl)
	}
	return &genericServer{
		srv:  srv,
		reqs: reqs,
		lck:  lck,
	}
}

func (g *genericServer) Server() *httptest.Server {
	return g.srv
}

func (g *genericServer) Requests() []*http.Request {
	g.lck.RLock()
	defer g.lck.RUnlock()
	return g.reqs
}
