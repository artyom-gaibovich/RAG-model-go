package config

import (
	"github.com/APSN4/RAG-model-go/app/controller"
	"github.com/APSN4/RAG-model-go/app/repository"
	"github.com/APSN4/RAG-model-go/app/service"
)

type Initialization struct {
	userRepo  repository.UserRepository
	userSvc   service.UserService
	UserCtrl  controller.UserController
	RoleRepo  repository.RoleRepository
	gptSvc    service.GPTService
	GPTCtrl   controller.GPTController
	LimitSvc  service.LimitService
	LimitCtrl controller.LimitController
}

func NewInitialization(userRepo repository.UserRepository,
	userService service.UserService,
	userCtrl controller.UserController,
	roleRepo repository.RoleRepository,
	gptSvc service.GPTService,
	GPTCtrl controller.GPTController,
	LimitSvc service.LimitService,
	LimitCtrl controller.LimitController) *Initialization {
	return &Initialization{
		userRepo:  userRepo,
		userSvc:   userService,
		UserCtrl:  userCtrl,
		RoleRepo:  roleRepo,
		gptSvc:    gptSvc,
		GPTCtrl:   GPTCtrl,
		LimitSvc:  LimitSvc,
		LimitCtrl: LimitCtrl,
	}
}
