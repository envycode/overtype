package repository

import (
	"context"
	"github.com/gofrs/uuid"
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
	GetRandom(ctx context.Context, contract contract.RequestGetContentContract) (schema.ContentTranslations, error)
	GetById(ctx context.Context, id uuid.UUID) (schema.ContentTranslations, error)
}

type ContentTranslationRepositoryImpl struct {
	Db *gorm.DB
}

func (c ContentTranslationRepositoryImpl) GetById(ctx context.Context, id uuid.UUID) (schema.ContentTranslations, error) {
	var contentTranslation schema.ContentTranslations
	res := c.Db.
		Model(&schema.ContentTranslations{}).
		Select("*").
		Where("content_translations.content_id = ?", id).
		First(&contentTranslation)

	if res.Error != nil {
		log.Errorln(res.Error)
		return schema.ContentTranslations{}, render.StatusError{
			HttpCode: http.StatusNotFound,
			Err:      res.Error,
		}
	}

	return contentTranslation, nil
}

func (c ContentTranslationRepositoryImpl) GetRandom(ctx context.Context, contract contract.RequestGetContentContract) (schema.ContentTranslations, error) {
	var contentTranslation schema.ContentTranslations
	res := c.Db.
		Model(&schema.ContentTranslations{}).
		Select("*").
		Where("content_translations.source_lang = ? AND content_translations.destined_lang = ?",
			contract.SourceLang,
			contract.DestinedLang).
		Order("RANDOM()").
		First(&contentTranslation)

	if res.Error != nil {
		log.Errorln(res.Error)
		return schema.ContentTranslations{}, render.StatusError{
			HttpCode: http.StatusNotFound,
			Err:      res.Error,
		}
	}

	return contentTranslation, nil
}

func (c ContentTranslationRepositoryImpl) Create(ctx context.Context, contract contract.RequestCreateContentContract) (schema.ContentTranslations, error) {
	contentDifficulty, err := schema.StrToDifficulty(contract.ContentDifficulty)
	if err != nil {
		log.Warningln(err)
		return schema.ContentTranslations{}, render.StatusError{
			HttpCode: http.StatusBadRequest,
			Err:      err,
		}
	}
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
		ContentDifficulty: contentDifficulty,
		SourceLang:        destinedLang,
		DestinedLang:      sourceLang,
		SourceText:        contract.SourceText,
		DestinedText:      contract.DestinedText,
		CreatedAt:         time.Time{},
		UpdatedAt:         time.Time{},
	}

	res := c.Db.Create(&contentTranslation)
	if res.Error != nil {
		log.Errorln(res.Error)
		return schema.ContentTranslations{}, res.Error
	}
	return contentTranslation, nil
}
