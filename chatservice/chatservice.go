package chatservice

import (
	"context"
	"log"

	"github.com/juandemanjon/go_service/gen"
)

// Server implement chat_service
type Server struct {
}

// NewServer makes a Server
func NewServer() *Server {
	return &Server{}
}

// SayHello rpc
func (s *Server) SayHello(ctx context.Context, in *gen.Message) (*gen.Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &gen.Message{Body: "Hello From the Server!"}, nil
}
