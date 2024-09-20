package controller

import (
	"context"
	"errors"

	"github.com/xpzouying/go-cmd-project-template/internal/model"
)

type User struct {
	Uid  int
	Name string
}

func GetUser(ctx context.Context, uid int) (*User, error) {

	user, found, err := model.GetUserByUid(ctx, uid)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, errors.New("user not found")
	}

	return &User{
		Uid:  int(user.ID),
		Name: user.Name,
	}, nil
}

func CreateUser(ctx context.Context, name string) (*User, error) {

	user := &model.User{
		Name: name,
	}
	if err := user.Create(ctx); err != nil {
		return nil, err
	}

	return &User{
		Uid:  int(user.ID),
		Name: user.Name,
	}, nil
}
