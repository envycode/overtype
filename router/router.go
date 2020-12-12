package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"overtype/appcontext"
	"overtype/handler"
	"overtype/repository"
	"overtype/service"
)

func SetupRoute(appCtx *appcontext.Application) *mux.Router {
	contentTranslationRepo := repository.ContentTranslationRepositoryImpl{
		Db: appCtx.Db,
	}

	addContentTranslationSvc := service.AddContentServiceImpl{
		Repo: contentTranslationRepo,
	}

	contentTranslationHandler := handler.ContentTranslationHandler{
		AppCtx:    appCtx,
		CreateSvc: addContentTranslationSvc,
	}

	r := mux.NewRouter()
	r.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("ok"))
	})

	privateRoute := r.PathPrefix("/internal").Subrouter()
	privateRoute.
		HandleFunc("/content-translations", contentTranslationHandler.Create).
		Methods(http.MethodPost)
	return r
}
