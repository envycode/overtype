package handler

import (
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
	"overtype/appcontext"
	"overtype/contract"
	"overtype/render"
	"overtype/service"
)

type PublicHandler struct {
	AppCtx                   *appcontext.Application
	GetTranslationContentSvc service.GetContentService
	CreateRoomSvc            service.CreateRoomService
	RoomSocketSvc            service.RoomSocketService
	GetContentByRoomSvc      service.GetContentByRoomService
}

func (p PublicHandler) Get(w http.ResponseWriter, r *http.Request) {
	getContentTranslationContract := contract.NewRequestGetContentContract(r)
	if err := p.AppCtx.Validate.Struct(getContentTranslationContract); err != nil {
		log.Infoln(err)
		render.RenderErr(w, render.StatusError{
			HttpCode: http.StatusBadRequest,
			Err:      err,
		})
		return
	}
	result, err := p.GetTranslationContentSvc.Exec(r.Context(), getContentTranslationContract)
	if err != nil {
		log.Infoln(err)
		render.RenderErr(w, err)
		return
	}
	render.RenderOk(w, result)
}

func (p PublicHandler) GetByRoomCode(w http.ResponseWriter, r *http.Request) {
	getContentTranslationContract := contract.NewRequestGetContentByRoomCodeContract(r)
	if err := p.AppCtx.Validate.Struct(getContentTranslationContract); err != nil {
		log.Infoln(err)
		render.RenderErr(w, render.StatusError{
			HttpCode: http.StatusBadRequest,
			Err:      err,
		})
		return
	}
	result, err := p.GetContentByRoomSvc.Exec(r.Context(), getContentTranslationContract)
	if err != nil {
		log.Infoln(err)
		render.RenderErr(w, err)
		return
	}
	render.RenderOk(w, result)
}

func (p PublicHandler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	createRoomContract := contract.NewRequestCreateRoomContract(r)
	if err := p.AppCtx.Validate.Struct(createRoomContract); err != nil {
		log.Infoln(err)
		render.RenderErr(w, render.StatusError{
			HttpCode: http.StatusBadRequest,
			Err:      err,
		})
		return
	}
	result, err := p.CreateRoomSvc.Exec(r.Context(), createRoomContract)
	if err != nil {
		log.Infoln(err)
		render.RenderErr(w, err)
		return
	}
	render.RenderOk(w, result)
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	EnableCompression: true,
}

func (p PublicHandler) JoinRoom(w http.ResponseWriter, r *http.Request) {
	joinRoomContract := contract.NewRequestJoinRoom(r)
	if err := p.AppCtx.Validate.Struct(joinRoomContract); err != nil {
		log.Infoln(err)
		render.RenderErr(w, render.StatusError{
			HttpCode: http.StatusBadRequest,
			Err:      err,
		})
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Errorln(err)
		render.RenderErr(w, err)
		return
	}

	defer func() {
		if err := conn.Close(); err != nil {
			log.Errorln(err)
			render.RenderErr(w, err)
		}
	}()

	wsWriterStream := make(chan string, 5)

	go p.RoomSocketSvc.Listener(r.Context(), joinRoomContract, conn, wsWriterStream)
	p.RoomSocketSvc.Writer(r.Context(), conn, wsWriterStream)
}
