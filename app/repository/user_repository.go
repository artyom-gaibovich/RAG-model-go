package repository

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"rag-model/app/domain/dao"
)

type UserRepository interface {
	FindAllUser() ([]dao.UserLogs, error)
	FindUserById(id int) (dao.UserLogs, error)
	Save(user *dao.UserLogs) (dao.UserLogs, error)
	DeleteUserById(id int) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (u UserRepositoryImpl) FindAllUser() ([]dao.UserLogs, error) {
	var users []dao.UserLogs

	var err = u.db.Preload("Role").Find(&users).Error
	if err != nil {
		log.Error("Got an error finding all couples. Error: ", err)
		return nil, err
	}

	return users, nil
}

func (u UserRepositoryImpl) FindUserById(id int) (dao.UserLogs, error) {
	user := dao.UserLogs{
		ID: id,
	}
	err := u.db.Preload("Role").First(&user).Error
	if err != nil {
		log.Error("Got and error when find user by id. Error: ", err)
		return dao.UserLogs{}, err
	}
	return user, nil
}

func (u UserRepositoryImpl) Save(user *dao.UserLogs) (dao.UserLogs, error) {
	var err = u.db.Save(user).Error
	if err != nil {
		log.Error("Got an error when save user. Error: ", err)
		return dao.UserLogs{}, err
	}
	return *user, nil
}

func (u UserRepositoryImpl) DeleteUserById(id int) error {
	err := u.db.Delete(&dao.UserLogs{}, id).Error
	if err != nil {
		log.Error("Got an error when delete user. Error: ", err)
		return err
	}
	return nil
}

func UserRepositoryInit(db *gorm.DB) *UserRepositoryImpl {
	db.AutoMigrate(&dao.UserLogs{})
	return &UserRepositoryImpl{
		db: db,
	}
}
