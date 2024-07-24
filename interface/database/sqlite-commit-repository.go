package database

import (
	"fmt"
	"log"

	"github.com/midedickson/github-service/dto"
	"gorm.io/gorm"
)

type SqliteCommitRepository struct {
	DB *gorm.DB
}

func NewSqliteCommitRepository(db *gorm.DB) *SqliteCommitRepository {
	return &SqliteCommitRepository{DB: db}
}

func (s *SqliteCommitRepository) StoreRepositoryCommits(commitRepoInfos *[]dto.CommitResponseDTO, repoName string, owner *User) error {
	//  logic to store commit info in the database
	repo := &Repository{}
	err := s.DB.Where("owner_id =?", owner.ID).Where("name =?", repoName).First(repo).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("repository not found for owner %v and repo %v", owner.Username, repoName)

		}
		return err
	}

	for _, commit := range *commitRepoInfos {
		// check if this commit already exists in our database
		existingCommit, err := s.GetCommitBySHA(commit.SHA)
		if err != nil {
			log.Println("Error in checking existing commits by sha")
			continue
		}
		if existingCommit != nil {
			// commit already exists, skip;
			log.Printf("Commit with SHA: %s already exists; skipping", existingCommit.SHA)
			continue
		}
		newCommit := &Commit{
			RepositoryName: repoName,
			SHA:            commit.SHA,
			Message:        commit.Message,
			Author:         commit.Author,
			Date:           commit.Date,
		}
		log.Printf("New commit to be created: %v", newCommit)
		err = s.DB.Create(newCommit).Error
		if err != nil {
			log.Printf("Error in saving commits with SHA: %s", newCommit.SHA)
			return err
		}
	}
	return nil
}

func (s *SqliteCommitRepository) GetCommitBySHA(sha string) (*Commit, error) {
	commit := &Commit{}
	err := s.DB.Where("sha =?", sha).First(commit).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return commit, nil
}

func (s *SqliteCommitRepository) GetRepositoryCommits(repoName string) ([]*Commit, error) {
	//  logic to retrieve commit info from the database by repository name
	commits := &[]*Commit{}
	err := s.DB.Where("repository_name =?", repoName).Find(commits).Error
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}
	return *commits, nil
}