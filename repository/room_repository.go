package repository

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"net/http"
	"overtype/contract"
	"overtype/render"
	"overtype/schema"
	"strings"
	"time"
)

type RoomRepository interface {
	Create(ctx context.Context, contract contract.RequestCreateRoomContract) (contract.ResponseCreateRoomContract, error)
	GetByCode(ctx context.Context, code string) (schema.Room, error)
	Join(ctx context.Context, in contract.RequestJoinRoomContract, room schema.Room) (schema.Room, error)
	Ready(ctx context.Context, state contract.RequestWebsocketContract, room schema.Room) (schema.Room, error)
	Sync(ctx context.Context, state contract.RequestWebsocketContract, room schema.Room) (schema.Room, error)
}

type RoomRepositoryImpl struct {
	Redis             *redis.Client
	ContentRepository ContentTranslationRepository
}

func (r RoomRepositoryImpl) Create(ctx context.Context, in contract.RequestCreateRoomContract) (contract.ResponseCreateRoomContract, error) {
	randomContentTranslation, err := r.ContentRepository.GetRandom(ctx, contract.RequestGetContentContract{
		SourceLang:   in.SourceLang,
		DestinedLang: in.DestinedLang,
	})

	if err != nil {
		log.Errorln(err)
		return contract.ResponseCreateRoomContract{}, err
	}

	roomId := randString(6)

	roomSpec := schema.Room{
		RoomId:  roomId,
		State:   schema.RoomStatePending,
		Content: randomContentTranslation,
		Result:  map[string]schema.RoomParticipant{},
	}

	if err := set(r.Redis, roomId, roomSpec); err != nil {
		log.Errorln(err)
		return contract.ResponseCreateRoomContract{}, err
	}

	return contract.ResponseCreateRoomContract{
		RoomId:  roomId,
		Created: true,
		Message: "room create success",
	}, nil
}

func (r RoomRepositoryImpl) GetByCode(ctx context.Context, code string) (schema.Room, error) {
	var room schema.Room

	if err := get(r.Redis, code, &room); err != nil {
		log.Errorln(err)
		return schema.Room{}, err
	}

	return room, nil
}

func (r RoomRepositoryImpl) Join(ctx context.Context, in contract.RequestJoinRoomContract, room schema.Room) (schema.Room, error) {
	if room.State == schema.RoomStateEnded {
		return schema.Room{}, render.StatusError{
			HttpCode: http.StatusBadRequest,
			Err:      errors.New("room already ended"),
		}
	}

	if room.State == schema.RoomStateFinish {
		return schema.Room{}, render.StatusError{
			HttpCode: http.StatusBadRequest,
			Err:      errors.New("room already finished"),
		}
	}

	room.Result[in.ParticipantId] = schema.RoomParticipant{
		State:            schema.ParticipantStatePending,
		CurrentWordCount: 0,
		ParticipantId:    in.ParticipantId,
		ParticipantName:  in.ParticipantName,
	}

	if err := set(r.Redis, room.RoomId, room); err != nil {
		log.Errorln(err)
		return schema.Room{}, err
	}

	return room, nil
}

func (r RoomRepositoryImpl) Ready(ctx context.Context, state contract.RequestWebsocketContract, room schema.Room) (schema.Room, error) {
	if room.State == schema.RoomStateEnded {
		return schema.Room{}, render.StatusError{
			HttpCode: http.StatusBadRequest,
			Err:      errors.New("room already ended"),
		}
	}

	if room.State == schema.RoomStateFinish {
		return schema.Room{}, render.StatusError{
			HttpCode: http.StatusBadRequest,
			Err:      errors.New("room already finished"),
		}
	}

	currentState := room.Result[state.ParticipantId]
	currentState.State = schema.ParticipantStateReady

	room.Result[state.ParticipantId] = currentState

	if err := set(r.Redis, room.RoomId, room); err != nil {
		log.Errorln(err)
		return schema.Room{}, err
	}

	isAllReady := true
	for _, v := range room.Result {
		if v.State != schema.ParticipantStateReady {
			isAllReady = false
			break
		}
	}

	if isAllReady {
	    room.State = schema.RoomStateStarted
        if err := set(r.Redis, room.RoomId, room); err != nil {
            log.Errorln(err)
            return schema.Room{}, err
        }
		res, err := json.Marshal(&contract.ResponseWebsocketContract{
			MyState:     currentState.State,
			RoomState:   schema.RoomStateStarted,
			LeaderBoard: room.Result,
		})
		if err != nil {
			log.Errorln("failed to marshal json: ", err)
		}
		r.Redis.Publish(room.RoomId, res)
	}

	return room, nil
}

func (r RoomRepositoryImpl) Sync(ctx context.Context, state contract.RequestWebsocketContract, room schema.Room) (schema.Room, error) {
	if room.State == schema.RoomStateEnded {
		return schema.Room{}, render.StatusError{
			HttpCode: http.StatusBadRequest,
			Err:      errors.New("room already ended"),
		}
	}

	if room.State == schema.RoomStateFinish {
		return schema.Room{}, render.StatusError{
			HttpCode: http.StatusBadRequest,
			Err:      errors.New("room already finished"),
		}
	}

	currentState := room.Result[state.ParticipantId]
	currentState.CurrentWordCount = state.CurrentWordCount

	wordLen := len(strings.Split(room.Content.DestinedText, " "))

	if currentState.CurrentWordCount >= wordLen {
		currentState.State = schema.ParticipantStateFinish
	}

	room.Result[state.ParticipantId] = currentState

	if err := set(r.Redis, room.RoomId, room); err != nil {
		log.Errorln(err)
		return schema.Room{}, err
	}

	isAllFinish := true
	for _, v := range room.Result {
		if v.State != schema.ParticipantStateFinish {
			isAllFinish = false
			break
		}
	}

	if isAllFinish {
		res, err := json.Marshal(&contract.ResponseWebsocketContract{
			MyState:     currentState.State,
			RoomState:   schema.RoomStateFinish,
			LeaderBoard: room.Result,
		})
		if err != nil {
			log.Errorln("failed to marshal json: ", err)
		}
		r.Redis.Publish(room.RoomId, res)
	}

	return room, nil
}

func set(c *redis.Client, key string, value interface{}) error {
	p, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.Set(key, p, time.Minute*10).Err()
}

func get(c *redis.Client, key string, dest interface{}) error {
	p := c.Get(key)
	if p.Err() != nil {
		return p.Err()
	}
	r, err := p.Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(r), dest)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
