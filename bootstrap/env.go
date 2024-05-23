package bootstrap

import (
	"send2kobo/logger"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Env struct {
	AppEnv                 string   `mapstructure:"APP_ENV"`
	ServerAddress          string   `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout         int      `mapstructure:"CONTEXT_TIMEOUT"`
	DBHost                 string   `mapstructure:"DB_HOST"`
	DBPort                 string   `mapstructure:"DB_PORT"`
	DBUser                 string   `mapstructure:"DB_USER"`
	DBPass                 string   `mapstructure:"DB_PASS"`
	DBName                 string   `mapstructure:"DB_NAME"`
	AccessTokenExpiryHour  int      `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int      `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string   `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string   `mapstructure:"REFRESH_TOKEN_SECRET"`
	UploadPath             string   `mapstructure:"UPLOAD_PATH"`
	KepubPath              string   `mapstructure:"KEPUB_PATH"`
	ZapLevel               string   `mapstructure:"ZAP_LEVEL"`
	ZapPath                []string `mapstructure:"ZAP_PATH"`
}

func NewEnv() *Env {
	env := &Env{}
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		logger.Error("Can't find the file .env : ", zap.Error(err))
	}
	err = viper.Unmarshal(env)
	if err != nil {
		logger.Error("Environment can't be loaded: ", zap.Error(err))
	}
	if env.AppEnv == "development" {
		logger.Infof("The App is running in development env")
	}
	return env
}
