package service

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
	"overtype/contract"
	"overtype/render"
	"overtype/repository"
	"overtype/schema"
	"time"
)

type RoomSocketService interface {
	Listener(ctx context.Context, req contract.RequestJoinRoomContract, conn *websocket.Conn, writer chan string)
	Writer(ctx context.Context, conn *websocket.Conn, writer chan string)
}

type RoomSocketServiceImpl struct {
	Redis *redis.Client
	Repo  repository.RoomRepository
}

func (r RoomSocketServiceImpl) Listener(ctx context.Context, req contract.RequestJoinRoomContract, conn *websocket.Conn, wsWriterStream chan string) {
	sub := r.Redis.Subscribe(req.Code)
	ch := sub.Channel()
	for {
		select {
		case message, ok := <-wsWriterStream:
			if !ok {
				return
			}
			if err := conn.SetWriteDeadline(time.Now().Add(10 * time.Second)); err != nil {
				log.Errorln("Fail receive message err:", err)
				continue
			}
			if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
				log.Errorln(err)
				continue
			}
		case broadcast := <-ch:
			wsWriterStream <- broadcast.Payload
		}

	}
}

func (r RoomSocketServiceImpl) Writer(ctx context.Context, conn *websocket.Conn, wsWriterStream chan string) {
	for {
		_, byteMsg, err := conn.ReadMessage()

		if err != nil {
			if !websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure) {
				log.Errorln("SocketClose err:", err)
			}
			return
		}

		var wsRequest contract.RequestWebsocketContract
		if err := json.Unmarshal(byteMsg, &wsRequest); err != nil {
			log.Errorln(err)
			r.sendAndMarshal(render.StatusError{
				HttpCode: http.StatusBadRequest,
				Err:      err,
			}, wsWriterStream)
			continue
		}

		r.Handle(ctx, wsRequest, wsWriterStream)
	}
}

func (r RoomSocketServiceImpl) Handle(ctx context.Context, req contract.RequestWebsocketContract, writer chan string) {
	room, err := r.Repo.GetByCode(ctx, req.Code)
	if err != nil {
		r.sendAndMarshal(render.StatusError{
			HttpCode: http.StatusBadRequest,
			Err:      err,
		}, writer)
		return
	}
	if req.Action == schema.ActionJoin {
		res, err := r.Repo.Join(ctx, contract.RequestJoinRoomContract{
			Code:          req.Code,
			ParticipantId: req.ParticipantId,
		}, room)
		if err != nil {
			r.sendAndMarshal(render.StatusError{
				HttpCode: http.StatusBadRequest,
				Err:      err,
			}, writer)
			return
		}
		r.sendAndMarshal(contract.ResponseWebsocketContract{
			MyState:     res.Result[req.ParticipantId].State,
			RoomState:   res.State,
			LeaderBoard: res.Result,
		}, writer)
		return
	}
	if req.Action == schema.ActionReady {
		res, err := r.Repo.Ready(ctx, req, room)
		if err != nil {
			r.sendAndMarshal(render.StatusError{
				HttpCode: http.StatusBadRequest,
				Err:      err,
			}, writer)
			return
		}
		if res.RoomId == "" {
			return
		}
		r.sendAndMarshal(contract.ResponseWebsocketContract{
			MyState:     res.Result[req.ParticipantId].State,
			RoomState:   res.State,
			LeaderBoard: res.Result,
		}, writer)
		return
	}
	if req.Action == schema.ActionSync {
		res, err := r.Repo.Sync(ctx, req, room)
		if err != nil {
			r.sendAndMarshal(render.StatusError{
				HttpCode: http.StatusBadRequest,
				Err:      err,
			}, writer)
			return
		}
		if res.RoomId == "" {
			return
		}
		r.sendAndMarshal(contract.ResponseWebsocketContract{
			MyState:     res.Result[req.ParticipantId].State,
			RoomState:   res.State,
			LeaderBoard: res.Result,
		}, writer)
		return
	}
}

func (r RoomSocketServiceImpl) sendAndMarshal(any interface{}, wsWriterStream chan string) {
	res, err := json.Marshal(&any)
	if err != nil {
		log.Errorln(err)
		wsWriterStream <- "failed to marshal json object from server"
		return
	}
	wsWriterStream <- string(res)
}
