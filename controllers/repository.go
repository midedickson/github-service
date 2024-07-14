package controllers

import (
	"net/http"

	"github.com/midedickson/github-service/utils"
)

func (c *Controller) GetRepositoryInfo(w http.ResponseWriter, r *http.Request) {
	owner, err := utils.GetPathParam(r, "owner")
	if err != nil {
		utils.Dispatch400Error(w, "Invalid Payload", err)
		return
	}
	user, err := c.dbRepository.GetUser(owner)
	if err != nil {
		utils.Dispatch500Error(w, err)
		return
	}
	if user == nil {
		utils.Dispatch404Error(w, "User with this github username not found, please register this github username", err)
		return
	}
	repoName, err := utils.GetPathParam(r, "repo")
	if err != nil {
		utils.Dispatch400Error(w, "Invalid Payload", err)
		return
	}
	repo, err := c.dbRepository.GetRepository(user.ID, repoName)
	if err != nil {
		utils.Dispatch500Error(w, err)
		return
	}
	if repo == nil {
		go c.task.AddRequestToFetchNewlyRequestedRepoQueue(user.Username, repoName)
		utils.Dispatch404Error(w, "Repository not found on Github; kindly check back again.", err)
		return
	}

	utils.Dispatch200(w, "Repository Information Fetched Successfully", repo)
}

func (c *Controller) GetRepositoryCommits(w http.ResponseWriter, r *http.Request) {
	repoName, err := utils.GetPathParam(r, "repo")
	if err != nil {
		utils.Dispatch400Error(w, "Invalid Payload", err)
		return
	}
	commits, err := c.dbRepository.GetRepositoryCommits(repoName)
	if err != nil {
		utils.Dispatch500Error(w, err)
		return
	}
	utils.Dispatch200(w, "Repository Commits Fetched Successfully", commits)
}
