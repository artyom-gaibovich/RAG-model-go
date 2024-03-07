package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"rag-model/app/constant"
	"rag-model/app/domain/dao"
	"rag-model/app/pkg"
	"rag-model/app/repository"
	"reflect"
	"strconv"
	"time"
)

type UserService interface {
	GetAllUser(c *gin.Context)
	GetUserById(c *gin.Context)
	AddUserData(c *gin.Context)
	UpdateUserData(c *gin.Context)
	UpdateComponentUserData(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func (u UserServiceImpl) UpdateUserData(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program update user data by id")
	userID, _ := strconv.Atoi(c.Param("userID"))

	var request dao.UserLogs
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := u.userRepository.FindUserById(userID)
	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	data.OperationType = request.OperationType
	data.Status = request.Status
	data.RoleID = request.RoleID
	data.UpdatedAt = time.Now()
	u.userRepository.Save(&data)

	if err != nil {
		log.Error("Happened error when updating data to database. Error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u UserServiceImpl) UpdateComponentUserData(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to update component current user by id")
	userID, _ := strconv.Atoi(c.Param("userID"))

	data, err := u.userRepository.FindUserById(userID)
	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	var request dao.UserLogs
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	fields := reflect.VisibleFields(reflect.TypeOf(request))
	for _, field := range fields {
		fmt.Printf("Key: %s\tType: %s\n", field.Name, field.Type)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u UserServiceImpl) GetUserById(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program get user by id")
	userID, _ := strconv.Atoi(c.Param("userID"))

	data, err := u.userRepository.FindUserById(userID)
	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u UserServiceImpl) AddUserData(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program add data user")
	var request dao.UserLogs
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	request.CreatedAt = time.Now()

	data, err := u.userRepository.Save(&request)
	if err != nil {
		log.Error("Happened error when saving data to database. Error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u UserServiceImpl) GetAllUser(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute get all data user")

	data, err := u.userRepository.FindAllUser()
	if err != nil {
		log.Error("Happened Error when find all user data. Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u UserServiceImpl) DeleteUser(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute delete data user by id")
	userID, _ := strconv.Atoi(c.Param("userID"))

	err := u.userRepository.DeleteUserById(userID)
	if err != nil {
		log.Error("Happened Error when try delete data user from DB. Error:", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null()))
}

func UserServiceInit(userRepository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}
