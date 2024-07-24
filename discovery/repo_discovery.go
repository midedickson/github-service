package discovery

import (
	"log"
	"sync"
	"time"

	"github.com/midedickson/github-service/dto"
	"github.com/midedickson/github-service/entity"
	"github.com/midedickson/github-service/interface/repository"
	"github.com/midedickson/github-service/requester"
)

type RepositoryDiscoveryService struct {
	requester        requester.Requester
	userRepository   repository.UserRepository
	repoRepository   repository.RepoRepository
	commitRepository repository.CommitRepository
}

func NewRepositoryDiscoveryService(requester requester.Requester,
	userRepository repository.UserRepository,
	repoRepository repository.RepoRepository,
	commitRepository repository.CommitRepository,
) *RepositoryDiscoveryService {
	return &RepositoryDiscoveryService{
		requester:        requester,
		userRepository:   userRepository,
		repoRepository:   repoRepository,
		commitRepository: commitRepository,
	}
}

func (rd *RepositoryDiscoveryService) GetAllUserRepositories(user *entity.User, wg *sync.WaitGroup) {
	//  logic to fetch all repositories for the given user
	defer wg.Done()
	// re-comfirm that this user is still in our database
	dbUser, _ := rd.userRepository.GetUser(user.Username)
	if dbUser == nil {
		log.Printf("User %v not found in database", user.Username)
		return
	}
	// Fetch all repositories for the user
	userRepositories, err := rd.requester.GetAllUserRepositories(user.Username)
	if err != nil {
		log.Printf("Error in fetching repositories for user %v: %v", user.Username, err)
		return
	}
	// using a go routine to optimize the saving of repositories and fetching the repo  commits
	// this will help the worker process tasks from the channel faster for users at scale
	for _, newRepoInfo := range *userRepositories {
		_, err := rd.repoRepository.StoreRepositoryInfo(&newRepoInfo, dbUser)
		if err != nil {
			log.Printf("Error in storing repository: %v", err)
			continue
		}
		log.Printf("fetching repository commits for repo: %s...", newRepoInfo.Name)
		remoteCommits, err := rd.requester.GetRepositoryCommits(user.Username, newRepoInfo.Name)
		if err != nil {
			log.Printf("Error in fetching commits: %v", err)
			continue
		}
		err = rd.commitRepository.StoreRepositoryCommits(remoteCommits, newRepoInfo.Name, dbUser)
		if err != nil {
			log.Printf("Error in saving commits: %v", err)
			continue
		}
	}
	log.Printf("Gotten repositories for user %v", user)
}

func (rd *RepositoryDiscoveryService) FetchNewlyRequestedRepo(repoRequest *dto.RepoRequest, wg *sync.WaitGroup) {
	//  logic to fetch a newly requested repo and commits for the given repository
	defer wg.Done()
	log.Println("waiting for newly requested repos...")

	remoteRepoInfo, err := rd.requester.GetRepositoryInfo(repoRequest.Username, repoRequest.RepoName)
	if err != nil {
		log.Printf("Error getting repository info: %v", err)
		return
	}
	user, _ := rd.userRepository.GetUser(repoRequest.Username)
	repo, _ := rd.repoRepository.StoreRepositoryInfo(remoteRepoInfo, user)
	log.Printf("fetching repository commits for repo: %s...", repoRequest.RepoName)
	remoteCommits, err := rd.requester.GetRepositoryCommits(user.Username, repo.Name)
	if err != nil {
		log.Printf("Error in fetching commits: %v", err)
		return
	}
	err = rd.commitRepository.StoreRepositoryCommits(remoteCommits, repo.Name, user)
	if err != nil {
		log.Printf("Error in saving commits: %v", err)
		return
	}
}

func (rd *RepositoryDiscoveryService) CheckForUpdateOnAllRepo() error {
	//  logic to check for updates on all repositories in the database

	allRepos, err := rd.repoRepository.GetAllRepositories()
	if err != nil {
		log.Printf("Error in fetching all repositories: %v", err)
		return err
	}

	for _, repo := range allRepos {

		log.Printf("Checking for updates on repo: %s...", repo.Name)
		remoteRepoInfo, err := rd.requester.GetRepositoryInfo(repo.Owner.Username, repo.Name)
		if err != nil {
			log.Printf("Error in fetching repository info: %v", err)
			continue
		}
		if repo.RemoteUpdatedAt != remoteRepoInfo.UpdatedAt {
			_, err = rd.repoRepository.StoreRepositoryInfo(remoteRepoInfo, repo.Owner)
			if err != nil {
				log.Printf("Error in updating repository: %v", err)
			}
		}
		// simulate more processing to reduce wasting ratelimit requests during tests
		time.Sleep(90 * time.Second)
	}
	return nil
}
