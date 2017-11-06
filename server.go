package main

import (
	"net/http"

	"github.com/go-chi/chi"
	com "github.com/utdta/website-backend/common"
)

func main() {
	// Validate environment
	com.ExitOnError(com.ValidateEnv())

	// Setup router
	r := chi.NewRouter()
	com.AttachMiddleware(r)

	r.Route("/api", func(r chi.Router) {
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			com.WriteJson(200, []byte(`HEALTH CHECK SUCCESS`), w)
		})
	})

	s := &http.Server{
		Addr:           `:` + com.ServerPort,
		Handler:        r,
		ReadTimeout:    com.REQUEST_READ_TIMEOUT_DURATION,
		WriteTimeout:   com.REQUEST_WRITE_TIMEOUT_DURATION,
		MaxHeaderBytes: com.REQUEST_MAX_HEADER_BYTES,
	}
	s.ListenAndServe()
}
