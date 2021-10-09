package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Server struct {
	client *mongo.Client
}

// NewServer is the constructor function for a Server struct
func NewServer() *Server {
	// create a new mongo client
  // TODO: Read db client from .env
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://nocandyonlywifi:qwertysucks@cluster0.d9vrx.mongodb.net/demo?retryWrites=true&w=majority")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
  
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

  return &Server{client}
}

// StartServer starts a new server at the given port
// taken as the function argument
// TODO: Read port from .env
func (s *Server) StartServer(port string) { 
  fmt.Printf("Starting a server on port: %s", port)
  if err := http.ListenAndServe(port, nil); err != nil {
    log.Fatalf("could not start server: %v", err)
  }
}
