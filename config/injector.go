// go:build wireinject
//go:build wireinject
// +build wireinject

package config

import (
	"github.com/APSN4/RAG-model-go/app/controller"
	"github.com/APSN4/RAG-model-go/app/repository"
	"github.com/APSN4/RAG-model-go/app/service"
	"github.com/google/wire"
)

var db = wire.NewSet(ConnectToDB)

var userServiceSet = wire.NewSet(service.UserServiceInit,
	wire.Bind(new(service.UserService), new(*service.UserServiceImpl)),
)

var userRepoSet = wire.NewSet(repository.UserRepositoryInit,
	wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
)

var userCtrlSet = wire.NewSet(controller.UserControllerInit,
	wire.Bind(new(controller.UserController), new(*controller.UserControllerImpl)),
)

var roleRepoSet = wire.NewSet(repository.RoleRepositoryInit,
	wire.Bind(new(repository.RoleRepository), new(*repository.RoleRepositoryImpl)),
)

var GPTServiceSet = wire.NewSet(service.GPTServiceInit,
	wire.Bind(new(service.GPTService), new(*service.GPTServiceImpl)),
)

var GPTCtrlSet = wire.NewSet(controller.GPTControllerInit,
	wire.Bind(new(controller.GPTController), new(*controller.GPTControllerImpl)),
)

var LimitServiceSet = wire.NewSet(service.LimitServiceInit,
	wire.Bind(new(service.LimitService), new(*service.LimitServiceImpl)),
)

var LimitCtrlSet = wire.NewSet(controller.LimitControllerInit,
	wire.Bind(new(controller.LimitController), new(*controller.LimitControllerImpl)),
)

func Init() *Initialization {
	wire.Build(NewInitialization, db, userCtrlSet, userServiceSet, userRepoSet, roleRepoSet, GPTCtrlSet, GPTServiceSet, LimitCtrlSet, LimitServiceSet)
	return nil
}
