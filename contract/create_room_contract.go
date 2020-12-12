package contract

import "net/http"

type RequestCreateRoomContract struct {
	SourceLang   string `json:"source_lang" validate:"required"`
	DestinedLang string `json:"destined_lang" validate:"required"`
}

func NewRequestCreateRoomContract(r *http.Request) RequestCreateRoomContract {
	params := r.URL.Query()
	return RequestCreateRoomContract{
		SourceLang:   params["source_lang"][0],
		DestinedLang: params["destined_lang"][0],
	}
}

type ResponseCreateRoomContract struct {
	RoomId  string `json:"room_id"`
	Created bool   `json:"created"`
	Message string `json:"message"`
}
