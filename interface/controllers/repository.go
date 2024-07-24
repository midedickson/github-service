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
	if owner == "" {
		utils.Dispatch400Error(w, "Invalid Payload", err)
		return
	}
	repoName, err := utils.GetPathParam(r, "repo")
	if err != nil {
		utils.Dispatch400Error(w, "Invalid Payload", err)
		return
	}
	if repoName == "" {
		utils.Dispatch400Error(w, "Invalid Payload", err)
		return
	}

	repo, err := c.repoUsecase.GetRepositoryInfo(owner, repoName)
	if err != nil {
		utils.Dispatch500Error(w, err)
		return
	}
	if repo == nil {
		utils.Dispatch404Error(w, "Repository not found on Github; kindly check back again.", err)
		return
	}
	utils.Dispatch200(w, "Repository Information Fetched Successfully", repo)
}

func (c *Controller) GetRepositories(w http.ResponseWriter, r *http.Request) {
	repoSearchParams := &utils.RepositorySearchParams{}
	owner, err := utils.GetPathParam(r, "owner")
	if err != nil || owner == "" {
		utils.Dispatch400Error(w, "Invalid Payload", err)
		return
	}
	utils.ParseQueryParams(r, repoSearchParams)
	repositories, err := c.repoUsecase.GetUserRepositories(owner, repoSearchParams)
	if err != nil {
		utils.Dispatch500Error(w, err)
		return
	}
	utils.Dispatch200(w, "Repositories Fetched Successfully", repositories)
}
