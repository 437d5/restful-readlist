// restful-readlist is a RESTful API with CRUD operations to db
// as a db PostgreSQL was chosen because it's the most popular relational database
// as a router Gorilla Mux was chosen
// server package contains connection to db and server creating
package main

import (
	"context"
	"fmt"
	"github.com/437d5/restful-readlist/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// it will stop everything if SIGINT or SIGTERM will be passed
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	connString := os.Getenv("CONN_STR")
	a := server.NewApp(ctx, connString)

	err := a.Start(ctx)
	if err != nil {
		fmt.Println("failed to start a new app:", err)
	}
}
