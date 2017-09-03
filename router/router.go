package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/vikashvverma/techscanservice/config"
	"github.com/vikashvverma/techscanservice/factory"
	"github.com/vikashvverma/techscanservice/handler"
	"github.com/vikashvverma/techscanservice/healthcheck"
)

const (
	GET  = "GET"
	POST = "POST"
)

func technologies(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{type:"PULL"}`))
}

func Router(c *config.Config, f *factory.Factory) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/healthcheck", healthcheck.Self).Methods(GET)
	router.HandleFunc("/api/techscan", handler.Technology(f.Fetcher(), f.Logger())).Methods(GET)
	router.HandleFunc("/api/techscan/{lang}", handler.Language(f.Fetcher(), f.Logger())).Methods(GET)
	router.HandleFunc("/api/techscan/{lang}/{page}", handler.Language(f.Fetcher(), f.Logger())).Methods(GET)
	router.HandleFunc("/api/owner/{repoID}", handler.Owner(f.Fetcher(), f.Logger())).Methods(GET)
	return router
}
