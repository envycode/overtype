package contract

import "overtype/schema"

type RequestWebsocketContract struct {
	Code             string `json:"code"`
	ParticipantId    string `json:"participant_id"`
	Action           string `json:"action"`
	CurrentWordCount int    `json:"current_word_count"`
}

type ResponseWebsocketContract struct {
	MyState     schema.ParticipantState           `json:"my_state"`
	RoomState   schema.RoomState                  `json:"room_state"`
	LeaderBoard map[string]schema.RoomParticipant `json:"leader_board"`
}
