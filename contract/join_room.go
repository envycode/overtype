package contract

import "net/http"

type RequestJoinRoomContract struct {
	Code            string `json:"code" validate:"required"`
	ParticipantId   string `json:"participant_id" validate:"required"`
	ParticipantName string `json:"participant_name" validate:"required"`
}

func NewRequestJoinRoom(r *http.Request) RequestJoinRoomContract {
	params := r.URL.Query()
	code := ""
	participantId := ""
	if len(params["code"]) > 0 {
		code = params["source_lang"][0]
	}
	if len(params["participant_id"]) > 0 {
		code = params["participant_id"][0]
	}
	return RequestJoinRoomContract{
		Code:          code,
		ParticipantId: participantId,
	}
}

type ResponseJoinRoomContract struct {
	Joined  bool   `json:"created"`
	Message string `json:"message"`
}
