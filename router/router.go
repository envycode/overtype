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

	getContentTranslationSvc := service.GetContentServiceImpl{
		Repo: contentTranslationRepo,
	}

	contentTranslationHandler := handler.ContentTranslationHandler{
		AppCtx:    appCtx,
		CreateSvc: addContentTranslationSvc,
	}

	publicHandler := handler.PublicHandler{
		AppCtx:                   appCtx,
		GetTranslationContentSvc: getContentTranslationSvc,
	}

	r := mux.NewRouter()
	r.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("ok"))
	})

	privateRoute := r.PathPrefix("/internal").Subrouter()
	privateRoute.
		HandleFunc("/content-translations", contentTranslationHandler.Create).
		Methods(http.MethodPost)

	apiRoute := r.PathPrefix("/api").Subrouter()
	apiRoute.
		HandleFunc("/content-translations", publicHandler.Get).
		Methods(http.MethodGet)

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./web_build")))
	return r
}
