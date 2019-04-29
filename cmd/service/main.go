package main

import (
	"net/http"
	"os"

	"github.com/chenpengcheng/boilerplates/service"
	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json"
	"github.com/segmentio/conf"
	"github.com/segmentio/events"
	_ "github.com/segmentio/events/text"
)

type Config struct {
	Addr   string `conf:"addr"`
	DBAddr string `conf:"db-addr"`
}

func main() {
	// Load configuration
	config := Config{
		Addr:   ":8080",
		DBAddr: "username:password@(localhost:3306)/product",
	}
	conf.Load(&config)
	events.Log("Started with %{config}v", config)

	// Create JSON-RPC 2.0 service
	s, err := service.New(service.Config{
		DBAddr: config.DBAddr,
	})
	if err != nil {
		events.Log("Stopped on %{error}v", err)
		os.Exit(1)
	}
	defer s.Close()

	// Register JSON-RPC 2.0 service
	rs := rpc.NewServer()
	rs.RegisterCodec(json.NewCodec(), "application/json")
	rs.RegisterService(s, service.Name)
	http.Handle("/rpc", rs)

	// Start JSON-RPC 2.0 service
	events.Log("Serving on %{address}s", config.Addr)
	events.Log("Stopped on %{error}v", http.ListenAndServe(config.Addr, nil))
}
