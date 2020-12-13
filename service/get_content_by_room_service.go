package service

import (
	"context"
	log "github.com/sirupsen/logrus"
	"overtype/contract"
	"overtype/repository"
)

type GetContentByRoomService interface {
	Exec(ctx context.Context, contract contract.RequestGetContentByRoomCodeContract) (contract.ResponseGetContentContract, error)
}

type GetContentByRoomServiceImpl struct {
	Repo repository.RoomRepository
}

func (a GetContentByRoomServiceImpl) Exec(ctx context.Context, data contract.RequestGetContentByRoomCodeContract) (contract.ResponseGetContentContract, error) {
	res, err := a.Repo.GetByCode(ctx, data.Code)
	if err != nil {
		log.Infoln("error getting content translations with spec: %v", res)
		return contract.ResponseGetContentContract{}, err
	}
	content := res.Content
	return contract.ResponseGetContentContract{
		ContentDifficulty: content.ContentDifficulty,
		SourceLang:        content.SourceLang,
		DestinedLang:      content.DestinedLang,
		SourceText:        content.SourceText,
		DestinedText:      content.DestinedText,
	}, nil
}
