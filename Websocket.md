initial connection
```
localhost:8080/api/join-room?code=XXX&participant_id=YYY&participant_name=ZZZ
```

websocket request contract
```
{
  "code": XXX,
  "participant_id": YYY,
  "action": "join" | "ready" | "sync",
  "current_word_count" : integer
}
```

websocket response contract
```
{
  "my_state": PARTICIPANT_STATE,
  "room_state": ROOM_STATE,
  "leader_board": key value participant state and word count
}
```

related code

```
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
```
