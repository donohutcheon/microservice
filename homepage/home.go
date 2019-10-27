package homepage

import (
	"log"
	"net/http"
	"time"
)

const message = "Hello World Microservice!"

//Handlers void
type Handlers struct {
	logger *log.Logger
}

//Home handles calls routed to "/""
func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	log.Println("HomeHandler received a request")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}

//Logger ---
func (h *Handlers) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer h.logger.Printf("request processed in %s\n", time.Now().Sub(startTime))
		next(w, r)
		
	}
}

//SetupRoutes add home route to mux
func (h *Handlers) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", h.Logger(h.Home))
}

//NewHandlers void
func NewHandlers(logger *log.Logger) *Handlers {
	return &Handlers{
		logger: logger,
	}

}
