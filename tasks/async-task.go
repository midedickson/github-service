package tasks

import (
	"github.com/midedickson/github-service/interface/database"
	"github.com/midedickson/github-service/interface/repository"
	"github.com/midedickson/github-service/requester"
)

type AsyncTask struct {
	GetAllRepoForUserQueue       chan *database.User
	FetchNewlyRequestedRepoQueue chan *RepoRequest
	CheckForUpdateOnAllRepoQueue chan string
	requester                    requester.Requester
	userRepository               repository.UserRepository
	repoRepository               repository.RepoRepository
	commitRepository             repository.CommitRepository
}

func NewAsyncTask(requester requester.Requester, dbRepository database.DBRepository, userRepository repository.UserRepository,
	repoRepository repository.RepoRepository,
	commitRepository repository.CommitRepository) *AsyncTask {
	return &AsyncTask{
		GetAllRepoForUserQueue:       make(chan *database.User),
		FetchNewlyRequestedRepoQueue: make(chan *RepoRequest),
		CheckForUpdateOnAllRepoQueue: make(chan string),
		requester:                    requester,
		userRepository:               userRepository,
		repoRepository:               repoRepository,
		commitRepository:             commitRepository,
	}
}
