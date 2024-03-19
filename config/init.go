package config

import (
	"rag-model/app/controller"
	"rag-model/app/repository"
	"rag-model/app/service"
)

type Initialization struct {
	userRepo repository.UserRepository
	userSvc  service.UserService
	UserCtrl controller.UserController
	RoleRepo repository.RoleRepository
	gptSvc   service.GPTService
	GPTCtrl  controller.GPTController
}

func NewInitialization(userRepo repository.UserRepository,
	userService service.UserService,
	userCtrl controller.UserController,
	roleRepo repository.RoleRepository,
	gptSvc service.GPTService,
	GPTCtrl controller.GPTController) *Initialization {
	return &Initialization{
		userRepo: userRepo,
		userSvc:  userService,
		UserCtrl: userCtrl,
		RoleRepo: roleRepo,
		gptSvc:   gptSvc,
		GPTCtrl:  GPTCtrl,
	}
}
