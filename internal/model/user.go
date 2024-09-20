package model

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name string `gorm:"column:name;type:varchar(255);not null"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) Create(ctx context.Context) error {
	return db.WithContext(ctx).Create(u).Error
}

func GetUserByUid(ctx context.Context, uid int) (*User, bool, error) {
	var user User
	if err := db.WithContext(ctx).
		Where("id = ?", uid).
		First(&user).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, false, nil
		}

		return nil, false, errors.Wrap(err, "failed to get user by uid")
	}

	return &user, true, nil

}
