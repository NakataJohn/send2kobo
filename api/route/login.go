package route

import (
	"send2kobo/api/controller"
	"send2kobo/bootstrap"
	"send2kobo/domain"
	"send2kobo/mongo"
	"send2kobo/repository"
	"send2kobo/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	lc := &controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur, timeout),
		Env:          env,
	}
	group.POST("/login", lc.Login)
	// group.GET("/login", lc.LoginPage)
}
