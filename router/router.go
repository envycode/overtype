package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"overtype/appcontext"
)

func SetupRoute(appCtx *appcontext.Application) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("ok"))
	})
	return r
}
