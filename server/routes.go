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

	r.HandleFunc("/add", urlHandler.AddHandler)
	r.HandleFunc("/get/{id:[0-9]+}", urlHandler.GetByIDHandler)
	r.HandleFunc("/get", urlHandler.GetHandler)
	r.HandleFunc("/update/{id:[0-9]+}", urlHandler.UpdateHandler)
	r.HandleFunc("/delete/{id:[0-9]+}", urlHandler.DeleteHandler)
}
