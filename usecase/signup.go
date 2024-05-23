package usecase

import (
	"context"
	"send2kobo/domain"
	"send2kobo/internal/tokenutil"
	"time"
)

type signupUsecase struct {
	userRepostory  domain.UserRepository
	contextTimeout time.Duration
}

func NewSignupUsecase(userRepostory domain.UserRepository, contextTimeout time.Duration) domain.SignupUsecase {
	return &signupUsecase{
		userRepostory:  userRepostory,
		contextTimeout: contextTimeout,
	}
}

func (su *signupUsecase) Create(c context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepostory.Create(ctx, user)
}

func (su *signupUsecase) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepostory.GetByEmail(ctx, email)
}

func (su *signupUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (su *signupUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}
