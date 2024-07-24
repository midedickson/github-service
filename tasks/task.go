package tasks

import (
	"github.com/midedickson/github-service/interface/database"
)

type Task interface {
	AddUserToGetAllRepoQueue(user *database.User)
	AddRequestToFetchNewlyRequestedRepoQueue(username, repoName string)
}
