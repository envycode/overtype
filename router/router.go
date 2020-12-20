package router

import (
	"fmt"
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

	roomRepository := repository.RoomRepositoryImpl{
		Redis:             appCtx.Redis,
		ContentRepository: contentTranslationRepo,
	}

	addContentTranslationSvc := service.AddContentServiceImpl{
		Repo: contentTranslationRepo,
	}

	getContentTranslationSvc := service.GetContentServiceImpl{
		Repo: contentTranslationRepo,
	}

	getContentByRoomCode := service.GetContentByRoomServiceImpl{
		Repo: roomRepository,
	}

	createRoomSvc := service.CreateRoomServiceImpl{
		Repo: roomRepository,
	}

	roomSocketSvc := service.RoomSocketServiceImpl{
		Redis: appCtx.Redis,
		Repo:  roomRepository,
	}

	contentTranslationHandler := handler.ContentTranslationHandler{
		AppCtx:    appCtx,
		CreateSvc: addContentTranslationSvc,
	}

	publicHandler := handler.PublicHandler{
		AppCtx:                   appCtx,
		GetTranslationContentSvc: getContentTranslationSvc,
		CreateRoomSvc:            createRoomSvc,
		RoomSocketSvc:            roomSocketSvc,
		GetContentByRoomSvc:      getContentByRoomCode,
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

	apiRoute.
		HandleFunc("/content-by-room-code", publicHandler.GetByRoomCode).
		Methods(http.MethodGet)

	apiRoute.
		HandleFunc("/create-room", publicHandler.CreateRoom).
		Methods(http.MethodPost)

	apiRoute.
		HandleFunc("/join-room", publicHandler.JoinRoom).
		Methods(http.MethodGet)

	fs := http.FileServer(http.Dir("./web_build"))

	r.PathPrefix("/room/{code}").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := mux.Vars(r)["code"]
		prefix := fmt.Sprintf("/room/%s", name)
		http.StripPrefix(prefix, fs).ServeHTTP(w, r)
	})

	r.PathPrefix("/").Handler(fs)
	return r
}
