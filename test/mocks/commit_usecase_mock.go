package mocks

import (
	"github.com/midedickson/github-service/entity"
	"github.com/stretchr/testify/mock"
)

type MockCommitUseCase struct {
	mock.Mock
}

func (m *MockCommitUseCase) GetRepositoryCommits(repoName string) ([]*entity.Commit, error) {
	args := m.Called(repoName)
	var commits []*entity.Commit
	if args.Get(0) != nil {
		commits = args.Get(0).([]*entity.Commit)
	}
	return commits, args.Error(1)
}
