package service

import (
	"context"
	log "github.com/sirupsen/logrus"
	"overtype/contract"
	"overtype/repository"
)

type GetContentService interface {
	Exec(ctx context.Context, contract contract.RequestGetContentContract) (contract.ResponseGetContentContract, error)
}

type GetContentServiceImpl struct {
	Repo repository.ContentTranslationRepository
}

func (a GetContentServiceImpl) Exec(ctx context.Context, data contract.RequestGetContentContract) (contract.ResponseGetContentContract, error) {
	res, err := a.Repo.GetRandom(ctx, data)
	if err != nil {
		log.Infoln("error getting content translations with spec: %v", res)
		return contract.ResponseGetContentContract{}, err
	}
	return contract.ResponseGetContentContract{
		ContentDifficulty:   res.ContentDifficulty,
		SourceLang:   res.SourceLang,
		DestinedLang: res.DestinedLang,
		SourceText:   res.SourceText,
		DestinedText: res.DestinedText,
	}, nil
}
