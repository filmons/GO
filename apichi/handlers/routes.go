package handlers

import (
    "github.com/go-chi/chi/v5"
    // "github.com/filmons/chiapi/db"
)

func Routes() *chi.Mux {
    r := chi.NewRouter()

    // Middleware example: logger
    r.Use(LoggingMiddleware)

    r.Get("/user/{userID}", GetUser) // GetUser is a handler defined in handlers.go

    return r
}
