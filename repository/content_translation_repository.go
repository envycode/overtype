package repository

import (
	"context"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"overtype/contract"
	"overtype/render"
	"overtype/schema"
	"time"
)

type ContentTranslationRepository interface {
	Create(ctx context.Context, contract contract.RequestCreateContentContract) (schema.ContentTranslations, error)
}

type ContentTranslationRepositoryImpl struct {
	Db *gorm.DB
}

func (c ContentTranslationRepositoryImpl) Create(ctx context.Context, contract contract.RequestCreateContentContract) (schema.ContentTranslations, error) {
	sourceLang, err := schema.StrToLang(contract.SourceLang)
	if err != nil {
		log.Warningln(err)
		return schema.ContentTranslations{}, render.StatusError{
			HttpCode: http.StatusBadRequest,
			Err:      err,
		}
	}
	destinedLang, err := schema.StrToLang(contract.DestinedLang)
	if err != nil {
		log.Warningln(err)
		return schema.ContentTranslations{}, render.StatusError{
			HttpCode: http.StatusBadRequest,
			Err:      err,
		}
	}
	contentTranslation := schema.ContentTranslations{
		SourceLang:   destinedLang,
		DestinedLang: sourceLang,
		SourceText:   contract.SourceText,
		DestinedText: contract.DestinedText,
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
	}

	res := c.Db.Create(&contentTranslation)
	if res.Error != nil {
		log.Errorln(res.Error)
		return schema.ContentTranslations{}, res.Error
	}
	return contentTranslation, nil
}
