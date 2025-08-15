package repo

import (
	"learn-go-goroutine/models"
	"learn-go-goroutine/types"

	"github.com/stretchr/testify/mock"
)

type userRepoMockDb struct {
	mock.Mock
}

func NewUserRepoMockDb() *userRepoMockDb {
	return &userRepoMockDb{}
}

func (u *userRepoMockDb) Create(user types.User) (models.User,error){
	args := u.Called(user)
	return args.Get(0).(models.User),args.Error(1)
}

func (u *userRepoMockDb) Update(user types.UpdateUser,id uint) error{
	args := u.Called(user,id)
	return args.Error(0)
}

func (u *userRepoMockDb) GetAll() (*[]models.User,error){
	args := u.Called()
	return args.Get(0).(*[]models.User),args.Error(1)
}