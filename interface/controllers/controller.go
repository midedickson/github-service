package controllers

import (
	"github.com/midedickson/github-service/interface/repository"
	"github.com/midedickson/github-service/requester"
	"github.com/midedickson/github-service/tasks"
)

type Controller struct {
	requester        requester.Requester
	userRepository   repository.UserRepository
	repoRepository   repository.RepoRepository
	commitRepository repository.CommitRepository
	task             tasks.Task
}

func NewController(
	requester requester.Requester,
	userRepository repository.UserRepository,
	repoRepository repository.RepoRepository,
	commitRepository repository.CommitRepository,
	task tasks.Task,
) *Controller {
	return &Controller{
		requester:        requester,
		userRepository:   userRepository,
		repoRepository:   repoRepository,
		commitRepository: commitRepository,
		task:             task,
	}
}
