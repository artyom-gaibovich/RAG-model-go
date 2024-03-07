package repository

import (
	"gorm.io/gorm"
	"rag-model/app/domain/dao"
)

type RoleRepository interface {
	FindAllRole()
}

type RoleRepositoryImpl struct {
	db *gorm.DB
}

func (r RoleRepositoryImpl) FindAllRole() {
	panic("implement me")
}

func RoleRepositoryInit(db *gorm.DB) *RoleRepositoryImpl {
	db.AutoMigrate(&dao.Role{}, &dao.UserLogs{})
	return &RoleRepositoryImpl{
		db: db,
	}
}
