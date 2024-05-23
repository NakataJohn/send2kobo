package route

import (
	"send2kobo/api/middleware"
	"send2kobo/bootstrap"
	"send2kobo/mongo"
	"time"

	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, db mongo.Database, gin *gin.Engine) {
	//必加才可以访问到
	// gin.LoadHTMLGlob("template/**/*")
	//加载静态资源
	// gin.Static("/assets", "./assets")

	gin.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	pubilcRouter := gin.Group("/api/v1")
	// pubilcRouter := gin.Group("")
	// All Public APIs
	timeout := time.Duration(env.ContextTimeout) * time.Second

	NewSignupRouter(env, timeout, db, pubilcRouter)
	NewLoginRouter(env, timeout, db, pubilcRouter)
	NewRefreshTokenRouter(env, timeout, db, pubilcRouter)

	protectedRouter := gin.Group("/api/v1")
	// protectedRouter := gin.Group("")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))

	// All Private APIs
	NewProfileRouter(env, timeout, db, protectedRouter)
	NewTaskRouter(env, timeout, db, protectedRouter)
	NewBookRouter(env, timeout, db, protectedRouter)
}
