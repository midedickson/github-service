package tasks

import (
	"github.com/midedickson/github-service/dto"
	"github.com/midedickson/github-service/entity"
	"github.com/midedickson/github-service/interface/discovery"
)

type TaskManager struct {
	GetAllRepoForUserQueue       chan *entity.User
	FetchNewlyRequestedRepoQueue chan *dto.RepoRequest
	CheckForUpdateOnAllRepoQueue chan string
	repoDiscovery                discovery.RepositoryDiscovery
}

func NewTaskManager(repoDiscovery discovery.RepositoryDiscovery) *TaskManager {
	return &TaskManager{
		GetAllRepoForUserQueue:       make(chan *entity.User),
		FetchNewlyRequestedRepoQueue: make(chan *dto.RepoRequest),
		CheckForUpdateOnAllRepoQueue: make(chan string),
		repoDiscovery:                repoDiscovery,
	}
}
