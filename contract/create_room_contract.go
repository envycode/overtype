package contract

import "net/http"

type RequestCreateRoomContract struct {
	SourceLang   string `json:"source_lang" validate:"required"`
	DestinedLang string `json:"destined_lang" validate:"required"`
}

func NewRequestCreateRoomContract(r *http.Request) RequestCreateRoomContract {
	params := r.URL.Query()
	sourceLang := ""
	destinedLang := ""
	if len(params["source_lang"]) > 0 {
		sourceLang = params["source_lang"][0]
	}
	if len(params["destined_lang"]) > 0 {
		destinedLang = params["destined_lang"][0]
	}
	return RequestCreateRoomContract{
		SourceLang:   sourceLang,
		DestinedLang: destinedLang,
	}
}

type ResponseCreateRoomContract struct {
	RoomId  string `json:"room_id"`
	Created bool   `json:"created"`
	Message string `json:"message"`
}
