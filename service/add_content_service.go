package service

import (
	"context"
	log "github.com/sirupsen/logrus"
	"overtype/contract"
	"overtype/repository"
)

type AddContentService interface {
	Exec(ctx context.Context, contract contract.RequestCreateContentContract) (contract.ResponseCreateContentContract, error)
}

type AddContentServiceImpl struct {
	Repo repository.ContentTranslationRepository
}

func (a AddContentServiceImpl) Exec(ctx context.Context, data contract.RequestCreateContentContract) (contract.ResponseCreateContentContract, error) {
	res, err := a.Repo.Create(ctx, data)
	if err != nil {
		log.Infoln("error creating content translations with spec: %v", res)
		return contract.ResponseCreateContentContract{}, err
	}
	return contract.ResponseCreateContentContract{
		Created: true,
		Message: "content translations succesfully created",
	}, nil
}
