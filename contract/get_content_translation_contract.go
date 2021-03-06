package contract

import (
	"net/http"
	"overtype/schema"
)

type RequestGetContentContract struct {
	ContentDifficulty string `json:"content_difficulty"`
	SourceLang        string `json:"source_lang" validate:"required"`
	DestinedLang      string `json:"destined_lang" validate:"required"`
}

func NewRequestGetContentContract(r *http.Request) RequestGetContentContract {
	params := r.URL.Query()
	contentDifficulty := ""
	sourceLang := ""
	destinedLang := ""
	if len(params["content_difficulty"]) > 0 {
		contentDifficulty = params["content_difficulty"][0]
	}
	if len(params["source_lang"]) > 0 {
		sourceLang = params["source_lang"][0]
	}
	if len(params["destined_lang"]) > 0 {
		destinedLang = params["destined_lang"][0]
	}
	return RequestGetContentContract{
		ContentDifficulty: contentDifficulty,
		SourceLang:        sourceLang,
		DestinedLang:      destinedLang,
	}
}

type RequestGetContentByRoomCodeContract struct {
	Code string `json:"code" validate:"required"`
}

func NewRequestGetContentByRoomCodeContract(r *http.Request) RequestGetContentByRoomCodeContract {
	params := r.URL.Query()
	code := ""
	if len(params["code"]) > 0 {
		code = params["code"][0]
	}
	return RequestGetContentByRoomCodeContract{
		Code: code,
	}
}

type ResponseGetContentContract struct {
	ContentDifficulty schema.ContentDifficulty `json:"content_difficulty" validate:"required"`
	SourceLang        schema.ContentLang       `json:"source_lang" validate:"required"`
	DestinedLang      schema.ContentLang       `json:"destined_lang" validate:"required"`
	SourceText        string                   `json:"source_text" validate:"required"`
	DestinedText      string                   `json:"destined_text" validate:"required"`
}
