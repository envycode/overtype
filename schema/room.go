package schema

const ActionJoin = "join"
const ActionReady = "ready"
const ActionSync = "sync"

type RoomState int

const RoomStatePending RoomState = 0
const RoomStateStarted RoomState = 1
const RoomStateFinish RoomState = 2
const RoomStateEnded RoomState = 3

type ParticipantState int

const ParticipantStatePending ParticipantState = 0
const ParticipantStateReady ParticipantState = 1
const ParticipantStateFinish ParticipantState = 2

type RoomParticipant struct {
	State            ParticipantState `json:"state"`
	CurrentWordCount int              `json:"word_type"`
}

type Room struct {
	RoomId  string                     `json:"room_id"`
	State   RoomState                  `json:"state"`
	Content ContentTranslations        `json:"content"`
	Result  map[string]RoomParticipant `json:"result"`
}
