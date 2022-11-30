package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jithinlal-gelato/go_api/handlers"
	"github.com/jithinlal-gelato/go_api/store"
	"log"
	"net/http"
)

type Args struct {
	conn string
	port string
}

func Run(args Args) error {
	router := mux.NewRouter().PathPrefix("/api/v1/").Subrouter()

	fmt.Println(args)

	st := store.NewPostgresEventStore(args.conn)
	hnd := handlers.NewEventHandler(st)
	RegisterAllRoutes(router, hnd)

	log.Println("Starting server at port: ", args.port)
	return http.ListenAndServe(args.port, router)
}

func RegisterAllRoutes(router *mux.Router, hnd handlers.IEventHandler) {
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})

	router.HandleFunc("/event", hnd.Get).Methods(http.MethodGet)
	// create events
	router.HandleFunc("/event", hnd.Create).Methods(http.MethodPost)
	// delete event
	router.HandleFunc("/event", hnd.Delete).Methods(http.MethodDelete)

	// cancel event
	router.HandleFunc("/event/cancel", hnd.Cancel).Methods(http.MethodPatch)
	// update event details
	router.HandleFunc("/event/details", hnd.UpdateDetails).Methods(http.MethodPut)
	// reschedule event
	router.HandleFunc("/event/reschedule", hnd.Reschedule).Methods(http.MethodPatch)

	// list events
	router.HandleFunc("/events", hnd.List).Methods(http.MethodGet)
}
