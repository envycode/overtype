package handler

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"overtype/appcontext"
	"overtype/contract"
	"overtype/render"
	"overtype/service"
)

type ContentTranslationHandler struct {
	AppCtx    *appcontext.Application
	CreateSvc service.AddContentService
}

func (c ContentTranslationHandler) Create(w http.ResponseWriter, r *http.Request) {
	createContentTranslationContract, err := contract.NewRequestCreateContentContract(r)
	if err != nil {
		log.Infoln(err)
		render.RenderErr(w, err)
		return
	}
	if err := c.AppCtx.Validate.Struct(createContentTranslationContract); err != nil {
		log.Infoln(err)
		render.RenderErr(w, render.StatusError{
			HttpCode: http.StatusBadRequest,
			Err:      err,
		})
		return
	}
	result, err := c.CreateSvc.Exec(r.Context(), createContentTranslationContract)
	if err != nil {
		log.Infoln(err)
		render.RenderErr(w, err)
		return
	}
	render.RenderOk(w, result)
}
