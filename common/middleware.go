package common

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	mw "github.com/pressly/chi/middleware"
)

// Hooks middleware onto router
func AttachMiddleware(r *chi.Mux) {
	r.Use(mw.RequestID)
	r.Use(mw.RealIP)
	r.Use(mw.Logger)
	r.Use(mw.Recoverer)
	r.Use(mw.CloseNotify)
	r.Use(mw.Timeout(MIDDLEWARE_TIMEOUT_DURATION))
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(cors.Handler)
}
