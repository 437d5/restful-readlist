package server

import (
	"github.com/437d5/restful-readlist/db"
	"github.com/437d5/restful-readlist/handlers"
	"github.com/gorilla/mux"
)

// newRoutes function adds router to App instance and loads it with endpoints
func (a *App) newRoutes() {
	r := mux.NewRouter()
	a.router = r

	a.loadRoutes(r)
}

func (a *App) loadRoutes(r *mux.Router) {
	urlHandler := &handlers.URLController{
		Controller: &db.Controller{
			Client: a.pgpool,
		},
	}

	r.HandleFunc("/", urlHandler.AddHandler).Methods("POST")
	r.HandleFunc("/{id:[0-9]+}", urlHandler.GetByIDHandler).Methods("GET")
	r.HandleFunc("/", urlHandler.GetHandler).Methods("GET")
	r.HandleFunc("/{id:[0-9]+}", urlHandler.UpdateHandler).Methods("PUT")
	r.HandleFunc("/{id:[0-9]+}", urlHandler.DeleteHandler).Methods("DELETE")
}
