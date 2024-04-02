// Package server response for opening new connection to db and creating a server
package server

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"net/http"
	"os"
	"time"
)

type App struct {
	router http.Handler
	pdb    *pgx.Conn
}

// NewApp creates new App entity including creating new connection to PostgresSQL
func NewApp() *App {
	// open a new connection to postgres
	pdb, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(fmt.Fprintf(os.Stderr, "Unable to create connection: %v\n", err))
	}
	defer func(pdb *pgx.Conn, ctx context.Context) {
		err := pdb.Close(ctx)
		if err != nil {
			log.Fatal(fmt.Fprintf(os.Stderr, "Unable to close connection: %v\n", err))
		}
	}(pdb, context.Background())

	a := &App{
		pdb: pdb,
	}
	a.newRoutes()
	return a
}

// Start function starts new server and Ping a connection to db to find out errors
// it also closes all connections and shutdowns a server
func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: a.router,

		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	err := a.pdb.Ping(ctx)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	defer func() {
		if err := a.pdb.Close(ctx); err != nil {
			fmt.Printf("failed to close connection to db: %v", err)
		}
	}()
	fmt.Println("Starting server")

	ch := make(chan error, 1)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to start server %w", err)
		}
		close(ch)
	}()

	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		return server.Shutdown(timeout)
	}
}
