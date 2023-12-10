package repository

import (
	"context"
	"entry-rpc/kitex_gen/user_service"
	"entry-rpc/modules"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByID(ctx context.Context, id int64) (*user_service.User, error)
	Save(ctx context.Context, user *user_service.User) error
	DeleteByID(ctx context.Context, id int64) error
}

func GetUserRepository() UserRepository {
	return &userRepository{
		mysqlDb: modules.MysqlCli,
	}
}

type userRepository struct {
	mysqlDb *gorm.DB
}

func (r *userRepository) FindByID(ctx context.Context, id int64) (*user_service.User, error) {
	var user user_service.User
	if err := r.mysqlDb.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Save(ctx context.Context, user *user_service.User) error {
	return r.mysqlDb.Save(user).Error
}
func (r *userRepository) DeleteByID(ctx context.Context, id int64) error {
	return r.mysqlDb.Model(&user_service.User{}).Delete("id", id).Error
}
