package appcontext

import (
	"github.com/go-playground/validator"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"overtype/config"
	"reflect"
	"strings"
	"time"
)

type Application struct {
	Config   config.Config
	Db       *gorm.DB
	Redis    *redis.Client
	Validate *validator.Validate
}

func NewApplication() *Application {
	validate := validator.New()
	useJsonFieldValidation(validate)
	cfg := config.NewConfig()
	return &Application{
		Config:   cfg,
		Db:       newDb(cfg),
		Redis:    newRedis(cfg),
		Validate: validate,
	}
}

func newRedis(config config.Config) *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:   config.Redis.Host,
		Dialer: nil,
	})

	if _, err := redisClient.Ping().Result(); err != nil {
		log.Fatalln(err)
	}
	return redisClient
}

func newDb(config config.Config) *gorm.DB {
	newLogger := logger.New(
		log.New(),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(postgres.Open(config.Db.Uri), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Fatalln(err)
	}
	return db
}

func useJsonFieldValidation(v *validator.Validate) {
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}
