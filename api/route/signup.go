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

func NewSignupRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	sc := controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(ur, timeout),
		Env:           env,
	}
	group.POST("/signup", sc.Signup)
	// group.GET("/signup", sc.SignupPage)
}
