package main

import (
	"log"
	"time"
)

type Server struct {
	host    string
	port    int
	timeout time.Duration
}

func (s *Server) Run() {
	log.Printf("Server running %s: %d", s.host, s.port)
}

// func newLocalHost() *Server {
// 	return &Server{
// 		host:    "127.0.0.1",
// 		port:    8080,
// 		timeout: 3 * time.Second,
// 	}
// }

// func newLocalHost(port interface{}, timeout interface{}) *Server {
// 	defaultPort := 8080
// 	defaultTimeout := 3 * time.Second

// 	actualPort := defaultPort
// 	if p, ok := port.(int); ok {
// 		actualPort = p
// 	}

// 	actualTimeout := defaultTimeout
// 	if t, ok := timeout.(time.Duration); ok {
// 		actualTimeout = t
// 	}

// 	return &Server{
// 		host:    "127.0.0.1",
// 		port:    actualPort,
// 		timeout: actualTimeout,
// 	}
// }

type OptionsServerFunc func(c *Server) error

func withTimeout(t time.Duration) OptionsServerFunc {
	return func(c *Server) error {
		c.timeout = t
		return nil
	}
}

func withPort(p int) OptionsServerFunc {
	return func(c *Server) error {
		c.port = p
		return nil
	}
}

func newLocalHost(opts ...OptionsServerFunc) (*Server, error) {
	server := &Server{
		host:    "127.0.0.1",
		port:    8080,
		timeout: 3 * time.Second,
	}

	for _, opt := range opts {
		if err := opt(server); err != nil {
			return nil, err
		}
	}
	return server, nil
}
func main() {
	localHostServer, err := newLocalHost(withTimeout(5*time.Second), withPort(7000))
	if err != nil {
		log.Fatal(err)
	}
	localHostServer.Run()
}
