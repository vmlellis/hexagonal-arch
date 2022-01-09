package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/vmlellis/go-hexagonal/application/adapters/web/server/handler"
	"github.com/vmlellis/go-hexagonal/application/contract"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

type Webserver struct {
	Service contract.ProductServiceInterface
}

func MakeNewWebserver(svc contract.ProductServiceInterface) *Webserver {
	return &Webserver{Service: svc}
}

func (w Webserver) Serve() {
	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)

	handler.MakeProducHandlers(r, n, w.Service)
	http.Handle("/", r)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":9000",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log:", log.Lshortfile),
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
