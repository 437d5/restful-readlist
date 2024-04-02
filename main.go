// restful-readlist is a RESTful API with CRUD operations to db
// as a db PostgreSQL was chosen because it's the most popular relational database
// as a router Gorilla Mux was chosen
// server package contains connection to db and server creating
package main

import (
	"context"
	"fmt"
	"github.com/437d5/restful-readlist/server"
	"os/signal"
	"syscall"
)

func main() {
	a := server.NewApp()

	// it will stop everything if SIGINT or SIGTERM will be passed
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	err := a.Start(ctx)
	if err != nil {
		fmt.Println("failed to start a new app:", err)
	}
}
