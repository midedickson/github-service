package tasks

import (
	"log"

	"github.com/midedickson/github-service/dto"
	"github.com/midedickson/github-service/entity"
)

func (t *TaskManager) AddUserToGetAllRepoQueue(user *entity.User) {
	t.GetAllRepoForUserQueue <- user
}

func (t *TaskManager) AddRequestToFetchNewlyRequestedRepoQueue(username, repoName string) {
	log.Println("Adding request to fetch newly requested")

	t.FetchNewlyRequestedRepoQueue <- &dto.RepoRequest{
		Username: username,
		RepoName: repoName,
	}
	log.Println("Added request to fetch newly requested")

}

func (t *TaskManager) AddSignalToCheckForUpdateOnAllRepoQueue() {
	t.CheckForUpdateOnAllRepoQueue <- "signal"
}
