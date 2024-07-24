package discovery

import (
	"sync"

	"github.com/midedickson/github-service/dto"
	"github.com/midedickson/github-service/entity"
)

type RepositoryDiscovery interface {
	GetAllUserRepositories(user *entity.User, wg *sync.WaitGroup)
	FetchNewlyRequestedRepo(repoRequest *dto.RepoRequest, wg *sync.WaitGroup)
	CheckForUpdateOnAllRepo() error
}
