package services

import (
	"github.com/PaulTabaco/bookstore_users-api/domain/users"
	"github.com/PaulTabaco/bookstore_users-api/utils/crypto_utils"
	"github.com/PaulTabaco/bookstore_users-api/utils/date_utils"
	"github.com/PaulTabaco/bookstore_utils/rest_errors"
)

type usersServiceInterface interface {
	GetUser(int64) (*users.User, rest_errors.RestErr)
	CreateUser(users.User) (*users.User, rest_errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, rest_errors.RestErr)
	DeleteUser(int64) rest_errors.RestErr
	SearchUser(string) (users.Users, rest_errors.RestErr)
	LoginUser(users.UserLoginRequest) (*users.User, rest_errors.RestErr) // =FindByEmailAndPassword
}

type usersService struct{}

var (
	// init UserService of type usersServiceInterface being INSTANCE of usersService
	UserService usersServiceInterface = &usersService{}
)

func (s *usersService) GetUser(userId int64) (*users.User, rest_errors.RestErr) {
	user := &users.User{Id: userId}

	if err := user.Get(); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *usersService) CreateUser(user users.User) (*users.User, rest_errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetNowDBFormat()
	user.Password = crypto_utils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, rest_errors.RestErr) {
	currentUser, err := s.GetUser(user.Id)

	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			currentUser.FirstName = user.FirstName
		}
		if user.LastName != "" {
			currentUser.LastName = user.LastName
		}
		if user.Email != "" {
			currentUser.Email = user.Email
		}
	} else {
		currentUser.FirstName = user.FirstName
		currentUser.LastName = user.LastName
		currentUser.Email = user.Email
	}

	if err := currentUser.Update(); err != nil {
		return nil, err
	}

	return currentUser, nil
}

func (s *usersService) DeleteUser(userId int64) rest_errors.RestErr {
	dao := &users.User{Id: userId}
	return dao.Delete()
}

func (s *usersService) SearchUser(status string) (users.Users, rest_errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}

func (s *usersService) LoginUser(request users.UserLoginRequest) (*users.User, rest_errors.RestErr) {
	dao := &users.User{
		Email:    request.Email,
		Password: crypto_utils.GetMd5(request.Password),
	}
	if err := dao.FindByEmailAndPassword(); err != nil {
		return nil, err
	}
	return dao, nil
}
