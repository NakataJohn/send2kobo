package domain

import "context"

type Profile struct {
	Name  string `json:"name"`  // 姓名
	Email string `json:"email"` // 邮箱
}

type ProfileUsecase interface {
	GetProfileByID(c context.Context, id string) (*Profile, error)
}
