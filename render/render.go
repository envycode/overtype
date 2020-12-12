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
	w.Header().Set("Content-Type", "application/json")
	jsondata, err := json.Marshal(data)
	if err != nil {
		RenderErr(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsondata)
}

func RenderErr(w http.ResponseWriter, err error) {
	switch e := err.(type) {
	case Error:
		w.WriteHeader(e.Status())
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(err.Error()))
}
