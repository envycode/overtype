package service

import (
	"context"
	"overtype/contract"
	"overtype/repository"
)

type CreateRoomService interface {
	Exec(ctx context.Context, req contract.RequestCreateRoomContract) (contract.ResponseCreateRoomContract, error)
}

type CreateRoomServiceImpl struct {
	Repo repository.RoomRepository
}

func (c CreateRoomServiceImpl) Exec(ctx context.Context, req contract.RequestCreateRoomContract) (contract.ResponseCreateRoomContract, error) {
	return c.Repo.Create(ctx, req)
}
