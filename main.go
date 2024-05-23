package main

import (
	"send2kobo/api/route"
	"send2kobo/bootstrap"
	_ "send2kobo/docs"
	"send2kobo/logger"

	"github.com/gin-gonic/gin"
)

// @title Send2Kobo Api
// @version v1.0.0
// @description Send2Kobo Api
// @termsOfService http://swagger.io/terms/
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey token
// @in header
// @name Authorization
func main() {
	// bootstrap.InitLogger()
	// defer bootstrap.Logger.Sync()
	logger.Infof("App is starting...")
	app := bootstrap.App()
	env := app.Env
	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	gin := gin.Default()

	route.Setup(env, db, gin)
	gin.Run(env.ServerAddress)
}
