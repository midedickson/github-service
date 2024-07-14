package database

import (
	"github.com/midedickson/github-service/dto"
	"github.com/midedickson/github-service/models"
)

type DBRepository interface {
	CreateUser(createUserPaylod *dto.CreateUserPayloadDTO) (*models.User, error)
	GetUser(username string) (*models.User, error)
	StoreRepositoryInfo(remoteRepoInfo *dto.RepositoryInfoResponseDTO, owner *models.User) (*models.Repository, error)
	GetRepository(ownerID uint, repoName string) (*models.Repository, error)
	StoreRepositoryCommits(commitRepoInfos *[]dto.CommitResponseDTO, repoName string, owner *models.User) error
	GetRepositoryCommits(repoName string) ([]*models.Commit, error)
	GetAllRepositories() ([]*models.Repository, error)
}
