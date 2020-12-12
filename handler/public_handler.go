package handler

import (
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
