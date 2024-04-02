package server

import (
	"github.com/437d5/restful-readlist/handlers"
	"github.com/gorilla/mux"
)

// newRoutes function adds router to App instance and loads it with endpoints
func (a *App) newRoutes() {
	r := mux.NewRouter()
	a.router = r

	r.HandleFunc("/add", handlers.AddHandler)
	r.HandleFunc("/get/{id:[0-9]+}", handlers.GetByIDHandler)
	r.HandleFunc("/get", handlers.GetHandler)
	r.HandleFunc("/update/{id:[0-9]+}", handlers.UpdateHandler)
	r.HandleFunc("/delete/{id:[0-9]+}", handlers.DeleteHandler)
}
