package discovery

import (
	"log"

	"github.com/midedickson/github-service/dto"
	"github.com/midedickson/github-service/entity"
	"github.com/midedickson/github-service/interface/repository"
	"github.com/midedickson/github-service/requester"
)

type CommitDiscoveryService struct {
	repoRepository   repository.RepoRepository
	commitRepository repository.CommitRepository
	requester        requester.Requester
}

func NewCommitDiscoveryService(repoRepository repository.RepoRepository,
	requester requester.Requester,
	commitRepository repository.CommitRepository) *CommitDiscoveryService {
	return &CommitDiscoveryService{
		repoRepository:   repoRepository,
		commitRepository: commitRepository,
		requester:        requester,
	}
}

func (cd *CommitDiscoveryService) GetLatestCommitSHAInRepository(repoName string) (string, error) {
	mostRecentCommit, err := cd.commitRepository.GetMostRecentCommitInRepository(repoName)
	if err != nil {
		return "", err
	}
	if mostRecentCommit != nil {
		return mostRecentCommit.SHA, nil
	}
	return "", nil
}

func (cd *CommitDiscoveryService) CheckForNewCommits(repo *entity.Repository) error {
	log.Printf("fetching new repository commits for repo: %s...", repo.Name)
	mostRecentSHA, err := cd.GetLatestCommitSHAInRepository(repo.Name)
	if err != nil {
		log.Printf("Error in fetching most recent commit SHA: %v", err)
		return err
	}
	newRemoteCommits, err := cd.requester.GetRepositoryCommits(repo.Owner.Username, repo.Name, &dto.CommitQueryParams{SHA: mostRecentSHA})
	if err != nil {
		log.Printf("Error in fetching new commits: %v", err)
		return err
	}
	err = cd.commitRepository.StoreRepositoryCommits(newRemoteCommits, repo.Name, repo.Owner)
	if err != nil {
		log.Printf("Error in saving new commits: %v", err)
		return err
	}
	return nil
}

func (cd *CommitDiscoveryService) GetCommitsForNewRepo(repo *entity.Repository) error {
	log.Printf("fetching repository commits for repo: %s...", repo.Name)
	remoteCommits, err := cd.requester.GetRepositoryCommits(repo.Owner.Username, repo.Name, nil)
	if err != nil {
		log.Printf("Error in fetching commits: %v", err)
		return err
	}
	err = cd.commitRepository.StoreRepositoryCommits(remoteCommits, repo.Name, repo.Owner)
	if err != nil {
		log.Printf("Error in saving commits: %v", err)
		return err
	}
	return nil
}

func (cd *CommitDiscoveryService) ResetCommitToSHA(repoName, resetSha string) error {
	log.Printf("resetting commits for repo: %s to SHA: %s...", repoName, resetSha)
	err := cd.commitRepository.DeleteUntilSHA(repoName, resetSha)
	if err != nil {
		log.Printf("Error in resetting commits: %v", err)
		return err
	}
	return nil
}
