package repository

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"overtype/contract"
	"overtype/schema"
	"time"
)

type RoomRepository interface {
	Create(ctx context.Context, contract contract.RequestCreateRoomContract) (contract.ResponseCreateRoomContract, error)
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
