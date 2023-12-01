package httpserver

import (
	"net"
	"time"
)

type Option func(*server)

func WithHostPort(host, port string) Option {
	return func(s *server) {
		s.server.Addr = net.JoinHostPort(host, port)
	}
}

func WithReadTimeout(timeout time.Duration) Option {
	return func(s *server) {
		s.server.ReadTimeout = timeout
	}
}

func WithWriteTimeout(timeout time.Duration) Option {
	return func(s *server) {
		s.server.WriteTimeout = timeout
	}
}

func WithShutdownTimeout(timeout time.Duration) Option {
	return func(s *server) {
		s.shutdownTimeout = timeout
	}
}
