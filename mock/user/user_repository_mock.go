package user

import (
	"CrowFundingV2/src/modules/user"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	Mock mock.Mock
}

func (repo *RepositoryMock) FindById(id string) *user.User {
	arguments := repo.Mock.Called(id)

	if arguments.Get(0) == nil {
		return nil
	}

	result := arguments.Get(0).(user.User)

	return &result
}
