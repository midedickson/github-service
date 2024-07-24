package usecase

import (
	"github.com/midedickson/github-service/entity"
	"github.com/midedickson/github-service/interface/repository"
)

type CommitUseCase interface {
	GetRepositoryCommits(repoName string) ([]*entity.Commit, error)
}

type CommitUseCaseService struct {
	commitRepository repository.CommitRepository
}

func NewCommitUseCaseService(commitRepository repository.CommitRepository) *CommitUseCaseService {
	return &CommitUseCaseService{commitRepository: commitRepository}
}

func (c *CommitUseCaseService) GetRepositoryCommits(repoName string) ([]*entity.Commit, error) {
	repoCommits, err := c.commitRepository.GetRepositoryCommits(repoName)
	if err != nil {
		return nil, err
	}
	commitEntities := make([]*entity.Commit, len(repoCommits))
	for i, commit := range repoCommits {
		commitEntities[i] = commit.ToEntity()
	}
	return commitEntities, nil
}
