package route

import (
	"send2kobo/bootstrap"
	"send2kobo/mongo"
	"time"

	"github.com/gin-gonic/gin"
)

func NewTaskRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	
	group.GET("/task")
	group.POST("/task")
}
