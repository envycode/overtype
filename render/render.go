package render

import (
	"encoding/json"
	"net/http"
)

type Error interface {
	error
	Status() int
}

type StatusError struct {
	HttpCode int
	Err      error
}

func (se StatusError) Error() string {
	return se.Err.Error()
}

func RenderOk(w http.ResponseWriter, data interface{}) {
	jsondata, err := json.Marshal(data)
	if err != nil {
		RenderErr(w, err)
		return
	}
	_, _ = w.Write(jsondata)
	w.WriteHeader(http.StatusOK)
}

func RenderErr(w http.ResponseWriter, err error) {
	switch e := err.(type) {
	case Error:
		w.WriteHeader(e.Status())
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	_, _ = w.Write([]byte(err.Error()))
}
