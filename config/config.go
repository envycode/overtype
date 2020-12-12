package config

import (
	"fmt"
	"github.com/subosito/gotenv"
	"log"
	"os"
)

type Config struct {
	App   App
	Db    Db
	Redis Redis
}

type App struct {
	Name string
	Host string
	Port int
	Env  string
}

type Db struct {
	Uri string
}

type Redis struct {
	Host string
}

func (a App) AppAddress() string {
	return fmt.Sprintf("%v:%v", a.Host, a.Port)
}

func NewConfig() Config {
	return Config{
		App: App{
			Name: GetString("APP_NAME"),
			Host: GetString("APP_HOST"),
			Port: GetInt("APP_PORT", 8080),
			Env:  GetString("APP_ENV"),
		},
		Db:    Db{Uri: GetString("DB_URI")},
		Redis: Redis{Host: GetString("REDIS_HOST")},
	}
}

var appConfig = Config{}

func init() {
	if err := gotenv.Load(); err != nil {
		log.Println("loading config from os environment variable")
	}
	log.SetOutput(os.Stdout)
	appConfig = NewConfig()
}

func AppConfig() Config {
	return appConfig
}
