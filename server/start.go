// Package server response for opening new connection to db and creating a server
package server

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"net/http"
	"os"
	"time"
)

type App struct {
	router http.Handler
	pgpool *pgxpool.Pool
}

// NewApp creates new App entity including creating new connection to PostgresSQL
func NewApp(ctx context.Context, connString string) *App {
	// open a new connection to postgres
	pgpool, err := pgxpool.New(ctx, connString)
	if err != nil {
		log.Fatal(fmt.Fprintf(os.Stderr, "Unable to create connection: %v\n", err))
	}
	defer pgpool.Close()

	a := &App{
		pgpool: pgpool,
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

	err := a.pgpool.Ping(ctx)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer a.pgpool.Close()

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
		timeout, cancel := context.WithTimeout(ctx, time.Second*10)
		defer cancel()
		return server.Shutdown(timeout)
	}
}
