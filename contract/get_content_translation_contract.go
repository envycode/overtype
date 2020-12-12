package contract

import (
	"net/http"
	"overtype/schema"
)

type RequestGetContentContract struct {
	SourceLang   string `json:"source_lang" validate:"required"`
	DestinedLang string `json:"destined_lang" validate:"required"`
}

func NewRequestGetContentContract(r *http.Request) RequestGetContentContract {
	params := r.URL.Query()
	return RequestGetContentContract{
		SourceLang:   params["source_lang"][0],
		DestinedLang: params["destined_lang"][0],
	}
}

type ResponseGetContentContract struct {
	SourceLang   schema.ContentLang `json:"source_lang" validate:"required"`
	DestinedLang schema.ContentLang `json:"destined_lang" validate:"required"`
	SourceText   string             `json:"source_text" validate:"required"`
	DestinedText string             `json:"destined_text" validate:"required"`
}
