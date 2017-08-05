package server

import (
	_ "expvar"
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/rms1000watt/degeneres-test/ballpark"
	"github.com/rms1000watt/degeneres-test/helpers"

	log "github.com/sirupsen/logrus"
)

func Ballpark(cfg Config, ballparkCfg ballpark.Config) {
	log.Debug("Ballpark Config: ", cfg)
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	if err := ballpark.PreServe(ballparkCfg); err != nil {
		log.Error("Failed running preserve function: ", err)
		return
	}

	srv := http.Server{
		Addr:              addr,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       30 * time.Second,
		Handler:           BallparkServerHandler(),
	}

	if cfg.CertsPath != "" && cfg.CertName != "" && cfg.KeyName != "" {
		log.Info("Starting HTTPS server at: ", addr)
		log.Fatal(srv.ListenAndServeTLS(filepath.Join(cfg.CertsPath, cfg.CertName), filepath.Join(cfg.CertsPath, cfg.KeyName)))
	} else {
		log.Info("Starting HTTP server at: ", addr)
		log.Fatal(srv.ListenAndServe())
	}
}

func BallparkServerHandler() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/debug/vars", http.DefaultServeMux)
	mux.HandleFunc("/person", helpers.HandleMiddlewares(personHandler, helpers.MiddlewareNoCache, helpers.MiddlewareCORS, helpers.MiddlewareLogger, helpers.MiddlewareSecure))
	mux.HandleFunc("/ticket", helpers.HandleMiddlewares(ticketHandler, helpers.MiddlewareNoCache, helpers.MiddlewareCORS, helpers.MiddlewareLogger, helpers.MiddlewareSecure))
	mux.HandleFunc("/management", helpers.HandleMiddlewares(managementHandler, helpers.MiddlewareNoCache, helpers.MiddlewareCORS, helpers.MiddlewareLogger, helpers.MiddlewareSecure))
	mux.HandleFunc("/", helpers.HandleMiddlewares(RootHandler, helpers.MiddlewareLogger))

	return mux
}

func personHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug("Starting personHandler")
	defer log.Debug("Finished personHandler")

	if r.Method == http.MethodOptions {
		// Process headers and return
		return
	}

	switch r.Method {
	case http.MethodGet:
		helpers.HandleMiddlewares(ballpark.PersonHandlerGET, helpers.MiddlewareNoCache)(w, r)
	case http.MethodPost:
		helpers.HandleMiddlewares(ballpark.PersonHandlerPOST, helpers.MiddlewareNoCache)(w, r)
	default:
		log.Debug("Method not allowed: ", r.Method)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

}

func ticketHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug("Starting ticketHandler")
	defer log.Debug("Finished ticketHandler")

	if r.Method == http.MethodOptions {
		// Process headers and return
		return
	}

	switch r.Method {
	case http.MethodGet:
		helpers.HandleMiddlewares(ballpark.TicketHandlerGET, helpers.MiddlewareNoCache)(w, r)
	case http.MethodPost:
		helpers.HandleMiddlewares(ballpark.TicketHandlerPOST, helpers.MiddlewareNoCache)(w, r)
	case http.MethodPut:
		helpers.HandleMiddlewares(ballpark.TicketHandlerPUT, helpers.MiddlewareNoCache)(w, r)
	default:
		log.Debug("Method not allowed: ", r.Method)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

}

func managementHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug("Starting managementHandler")
	defer log.Debug("Finished managementHandler")

	if r.Method == http.MethodOptions {
		// Process headers and return
		return
	}

	helpers.HandleMiddlewares(ballpark.ManagementHandler, helpers.MiddlewareNoCache)(w, r)

}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug("Path not found: ", r.URL.Path)
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}
