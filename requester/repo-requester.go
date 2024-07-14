package requester

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/midedickson/github-service/dto"
	"github.com/midedickson/github-service/utils"
)

type RepositoryRequester struct {
	http.Client
}

func NewRepositoryRequester() *RepositoryRequester {
	return &RepositoryRequester{}
}

func (r *RepositoryRequester) GetRepositoryInfo(owner, repo string) (*dto.RepositoryInfoResponseDTO, error) {
	// Implement logic to fetch repository info
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s", owner, repo)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return nil, utils.ErrRepoNotFound
	}

	var repository dto.RepositoryInfoResponseDTO
	if err := json.NewDecoder(resp.Body).Decode(&repository); err != nil {
		return nil, err
	}
	return &repository, nil
}
func (r *RepositoryRequester) GetRepositoryCommits(owner, repo string) (*[]dto.CommitResponseDTO, error) {
	// Implement logic to fetch repository commits
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/commits", owner, repo)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var commits []dto.CommitResponseDTO
	if err := json.NewDecoder(resp.Body).Decode(&commits); err != nil {
		return nil, err
	}
	return &commits, nil
}

func (r *RepositoryRequester) GetAllUserRepositories(owner string) (*[]dto.RepositoryInfoResponseDTO, error) {
	// Implement logic to fetch all repositories for a user
	url := fmt.Sprintf("https://api.github.com/users/%s/repos", owner)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var repositories []dto.RepositoryInfoResponseDTO
	if err := json.NewDecoder(resp.Body).Decode(&repositories); err != nil {
		return nil, err
	}
	return &repositories, nil
}
