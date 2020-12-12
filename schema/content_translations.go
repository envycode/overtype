package schema

import (
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type ContentLang string

const LangEn ContentLang = "en"
const LangJp ContentLang = "jp"
const LangJpHiragana ContentLang = "jp-hiragana"
const LangJpKatakana ContentLang = "jp-katakana"

type ContentTranslations struct {
	ContentId    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	SourceLang   ContentLang
	DestinedLang ContentLang
	SourceText   string
	DestinedText string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (c *ContentTranslations) BeforeCreate(tx *gorm.DB) (err error) {
	c.ContentId = uuid.NewV4()
	return
}

func StrToLang(in string) (ContentLang, error) {
	switch in {
	case "jp":
		return LangJp, nil
	case "en":
		return LangEn, nil
	case "jp-hiragana":
		return LangJpHiragana, nil
	case "jp-katakana":
		return LangJpKatakana, nil
	default:
		return "", errors.New(fmt.Sprintf("lang: %v not found"))
	}
}
