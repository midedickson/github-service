package controllers

import (
	"log"
	"net/http"

	"github.com/midedickson/github-service/utils"
)

func (c *Controller) GetRepositoryCommits(w http.ResponseWriter, r *http.Request) {
	repoName, err := utils.GetPathParam(r, "repo")
	if err != nil {
		utils.Dispatch400Error(w, "Invalid Payload", err)
		return
	}
	if repoName == "" {
		utils.Dispatch400Error(w, "Invalid Payload", err)
		return
	}
	commits, err := c.commitRepository.GetRepositoryCommits(repoName)
	if err != nil {
		log.Printf("%v", err)
		utils.Dispatch500Error(w, err)
		return
	}
	utils.Dispatch200(w, "Repository Commits Fetched Successfully", commits)
}
