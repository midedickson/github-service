package controllers_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/midedickson/github-service/entity"
	"github.com/midedickson/github-service/interface/controllers"
	"github.com/midedickson/github-service/test/mocks"
	"github.com/midedickson/github-service/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetRepositoryCommits(t *testing.T) {
	mockCommitUseCase := new(mocks.MockCommitUseCase)
	controller := controllers.NewController(nil, nil, nil, mockCommitUseCase)

	t.Run("successful fetch repository commits", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/{owner}/repos/{repo}/commits", nil)
		assert.NoError(t, err)
		req = mux.SetURLVars(req, map[string]string{"owner": "testuserx", "repo": "testrepo"})

		rr := httptest.NewRecorder()
		commits := []*entity.Commit{
			{Message: "Initial commit", Author: "testuserx"},
			{Message: "Added new feature", Author: "testuserx"},
		}
		mockCommitUseCase.On("GetRepositoryCommits", "testrepo").Return(commits, nil)

		http.HandlerFunc(controller.GetRepositoryCommits).ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		var response utils.APIResponse
		json.Unmarshal(rr.Body.Bytes(), &response)
		assert.Equal(t, true, response.Success)
		assert.Equal(t, "Repository Commits Fetched Successfully", response.Message)
		mockCommitUseCase.AssertExpectations(t)
	})

	t.Run("invalid payload - missing repo", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/{owner}/repos//commits", nil)
		assert.NoError(t, err)
		req = mux.SetURLVars(req, map[string]string{"owner": "testuserx"})

		rr := httptest.NewRecorder()

		http.HandlerFunc(controller.GetRepositoryCommits).ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		var response utils.APIResponse
		json.Unmarshal(rr.Body.Bytes(), &response)
		assert.Equal(t, false, response.Success)
		assert.Equal(t, "Invalid Payload", response.Message)
	})

	t.Run("internal server error", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/{owner}/repos/{repo}/commits", nil)
		assert.NoError(t, err)
		req = mux.SetURLVars(req, map[string]string{"owner": "testuserx", "repo": "testrepox"})

		rr := httptest.NewRecorder()

		mockCommitUseCase.On("GetRepositoryCommits", "testrepox").Return(nil, errors.New("some error"))

		http.HandlerFunc(controller.GetRepositoryCommits).ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		var response utils.APIResponse
		json.Unmarshal(rr.Body.Bytes(), &response)
		assert.Equal(t, false, response.Success)
		assert.Equal(t, "some error", response.Message)
		mockCommitUseCase.AssertExpectations(t)
	})
}
