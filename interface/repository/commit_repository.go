package repository

import (
	"github.com/midedickson/github-service/dto"
	"github.com/midedickson/github-service/interface/database"
)

type CommitRepository interface {
	StoreRepositoryCommits(commitRepoInfos *[]dto.CommitResponseDTO, repoName string, owner *database.User) error
	GetRepositoryCommits(repoName string) ([]*database.Commit, error)
}
