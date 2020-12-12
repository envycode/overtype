package contract

import (
	"encoding/json"
	"net/http"
	"overtype/render"
)

type RequestCreateContentContract struct {
	ContentDifficulty string `json:"content_difficulty" validate:"required"`
	SourceLang        string `json:"source_lang" validate:"required"`
	DestinedLang      string `json:"destined_lang" validate:"required"`
	SourceText        string `json:"source_text" validate:"required"`
	DestinedText      string `json:"destined_text" validate:"required"`
}

func NewRequestCreateContentContract(r *http.Request) (RequestCreateContentContract, error) {
	var createContract RequestCreateContentContract
	err := json.NewDecoder(r.Body).Decode(&createContract)
	if err != nil {
		return RequestCreateContentContract{}, render.StatusError{
			HttpCode: http.StatusBadRequest,
			Err:      err,
		}
	}
	return createContract, nil
}

type ResponseCreateContentContract struct {
	Created bool   `json:"created"`
	Message string `json:"message"`
}
