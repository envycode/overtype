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
	participantName := "Anonymous"
	if len(params["code"]) > 0 {
		code = params["code"][0]
	}
	if len(params["participant_id"]) > 0 {
		participantId = params["participant_id"][0]
	}
	if len(params["participant_name"]) > 0 {
		participantName = params["participant_name"][0]
	}
	return RequestJoinRoomContract{
		Code:            code,
		ParticipantId:   participantId,
		ParticipantName: participantName,
	}
}

type ResponseJoinRoomContract struct {
	Joined  bool   `json:"created"`
	Message string `json:"message"`
}
