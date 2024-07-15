package controllers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/gorilla/mux"
	"github.com/midedickson/github-service/controllers"
	"github.com/midedickson/github-service/mocks"
	"github.com/midedickson/github-service/models"
	"github.com/midedickson/github-service/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetRepositoryCommits(t *testing.T) {
	// Initialize the mocks
	mockDBRepository := new(mocks.MockDBRepository)
	mockRequester := new(mocks.MockRequester)
	mockTask := new(mocks.MockTask)

	// Create the controller with mocked dependencies
	controller := controllers.NewController(mockRequester, mockDBRepository, mockTask)

	// Test cases
	tests := []struct {
		name          string
		repoName      string
		mockSetup     func()
		expectedCode  int
		expectedError string
	}{
		{
			name:          "Invalid repo path parameter",
			repoName:      "",
			mockSetup:     func() {},
			expectedCode:  http.StatusBadRequest,
			expectedError: "Invalid Payload",
		},
		{
			name:     "Database error while fetching commits",
			repoName: "testrepo",
			mockSetup: func() {
				mockDBRepository.On("GetRepositoryCommits", "testrepo").Return([]*models.Commit{}, assert.AnError)
			},
			expectedCode:  http.StatusInternalServerError,
			expectedError: "assert.AnError general error for testing",
		},
		{
			name:     "Successful fetch of repository commits",
			repoName: "testrepox",
			mockSetup: func() {
				commits := []*models.Commit{
					{SHA: "commitsha", Message: "commit message", Author: "author", Date: "date"},
				}
				mockDBRepository.On("GetRepositoryCommits", "testrepox").Return(commits, nil)
			},
			expectedCode:  http.StatusOK,
			expectedError: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up the mocks
			tt.mockSetup()

			// Create a new HTTP request
			req, _ := http.NewRequest("GET", "/repos/{repo}/commits", nil)
			rr := httptest.NewRecorder()
			req = mux.SetURLVars(req, map[string]string{"repo": tt.repoName})

			// Call the GetRepositoryCommits method
			controller.GetRepositoryCommits(rr, req)

			// Check the response status code and body
			assert.Equal(t, tt.expectedCode, rr.Code)
			if tt.expectedError != "" {
				var response utils.APIResponse
				json.Unmarshal(rr.Body.Bytes(), &response)
				assert.Equal(t, false, response.Success)
				assert.Equal(t, tt.expectedError, response.Message)
			} else {
				var response utils.APIResponse
				json.Unmarshal(rr.Body.Bytes(), &response)
				assert.Equal(t, true, response.Success)
				assert.Equal(t, "Repository Commits Fetched Successfully", response.Message)
			}

			// Assert that the expectations were met
			mockDBRepository.AssertExpectations(t)
		})
	}
}

func TestGetRepositoryInfo(t *testing.T) {
	// Initialize the mocks
	mockDBRepository := new(mocks.MockDBRepository)
	mockRequester := new(mocks.MockRequester)
	mockTask := new(mocks.MockTask)

	// Create the controller with mocked dependencies
	controller := controllers.NewController(mockRequester, mockDBRepository, mockTask)

	// Test cases
	tests := []struct {
		name          string
		owner         string
		repoName      string
		mockSetup     func()
		expectedCode  int
		expectedError string
	}{
		{
			name:          "Invalid owner path parameter",
			owner:         "",
			repoName:      "testrepo",
			mockSetup:     func() {},
			expectedCode:  http.StatusBadRequest,
			expectedError: "Invalid Payload",
		},
		{
			name:     "User not found in database",
			owner:    "testuser",
			repoName: "testrepo",
			mockSetup: func() {
				mockDBRepository.On("GetUser", "testuser").Return(nil, nil)
			},
			expectedCode:  http.StatusNotFound,
			expectedError: "User with this github username not found, please register this github username",
		},
		{
			name:     "Database error while fetching user",
			owner:    "testuser",
			repoName: "testrepo",
			mockSetup: func() {
				mockDBRepository.On("GetUser", "testuser").Return(nil, assert.AnError)
			},
			expectedCode:  http.StatusInternalServerError,
			expectedError: "assert.AnError",
		},
		{
			name:          "Invalid repo path parameter",
			owner:         "testuser",
			repoName:      "",
			mockSetup:     func() {},
			expectedCode:  http.StatusBadRequest,
			expectedError: "Invalid Payload",
		},
		{
			name:     "Repository not found in database",
			owner:    "testuser",
			repoName: "testrepo",
			mockSetup: func() {
				user := &models.User{Username: "testuser"}
				mockDBRepository.On("GetUser", "testuser").Return(user, nil)
				mockDBRepository.On("GetRepository", user.ID, "testrepo").Return(nil, nil)
				mockTask.On("AddRequestToFetchNewlyRequestedRepoQueue", "testuser", "testrepo").Return()
			},
			expectedCode:  http.StatusNotFound,
			expectedError: "Repository not found on Github; kindly check back again.",
		},
		{
			name:     "Database error while fetching repository",
			owner:    "testuser",
			repoName: "testrepo",
			mockSetup: func() {
				user := &models.User{Username: "testuser"}
				mockDBRepository.On("GetUser", "testuser").Return(user, nil)
				mockDBRepository.On("GetRepository", user.ID, "testrepo").Return(nil, assert.AnError)
			},
			expectedCode:  http.StatusInternalServerError,
			expectedError: "assert.AnError",
		},
		{
			name:     "Successful fetch of repository information",
			owner:    "testuser",
			repoName: "testrepo",
			mockSetup: func() {
				user := &models.User{Username: "testuser"}
				repo := &models.Repository{Name: "testrepo"}
				mockDBRepository.On("GetUser", "testuser").Return(user, nil)
				mockDBRepository.On("GetRepository", user.ID, "testrepo").Return(repo, nil)
			},
			expectedCode:  http.StatusOK,
			expectedError: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Skip()
			// Set up the mocks
			tt.mockSetup()

			// Create a new HTTP request
			req, _ := http.NewRequest("GET", "/repos/{owner}/{repo}", nil)
			rr := httptest.NewRecorder()
			req = mux.SetURLVars(req, map[string]string{"owner": tt.owner, "repo": tt.repoName})

			var wg sync.WaitGroup
			wg.Add(1)
			mockTask.On("AddRequestToFetchNewlyRequestedRepoQueue", tt.owner, tt.repoName).Run(func(args mock.Arguments) {
				wg.Done()
			}).Return()

			// Call the GetRepositoryInfo method
			controller.GetRepositoryInfo(rr, req)
			wg.Wait()

			// Check the response status code and body
			assert.Equal(t, tt.expectedCode, rr.Code)
			if tt.expectedError != "" {
				var response utils.APIResponse
				json.Unmarshal(rr.Body.Bytes(), &response)
				assert.Equal(t, false, response.Success)
				assert.Equal(t, tt.expectedError, response.Message)
			}

			// Assert that the expectations were met
			mockDBRepository.AssertExpectations(t)
			mockTask.AssertExpectations(t)
		})
	}
}
