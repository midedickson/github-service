# GitHub Service

## Overview

GitHub Service is a Go-based application that interacts with the GitHub API to manage repositories and their commits. The service allows users to register their GitHub usernames, fetch repository information, and retrieve commit details for repositories.

## Table of Contents

1. [Getting Started](#getting-started)
2. [Project Structure](#project-structure)
3. [Dependencies](#dependencies)
4. [Usage](#usage)
5. [Running Tests](#running-tests)
6. [API Endpoints](#api-endpoints)

## Getting Started

### Prerequisites

Before you begin, ensure you have the following installed on your machine:

- [Go](https://golang.org/doc/install) (version 1.15+)
- [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

### Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/midedickson/github-service.git
   ```

2. Navigate to the project directory:

   ```sh
   cd github-service
   ```

3. Download the project dependencies:

   ```sh
   go mod download
   ```

## Project Structure

```
- github-service/
  - cmd/
    - main.go
  - config/
    - env.go
  - discovery/
    - interface.go
    - repo_discovery.go
  - dto/
    - create_user_payload.go
    - api_response.go
  - entity/
    - repository.go
  - interface/
    - controllers/
      - commits.go
      - commits_test.go
    - database/
     - commit.go
     - db.go
     - repository.go
     - sqlite-commit-repository.go
     - sqlite-repo-repository.go
     - sqlite-user-repository.go
     - user.go
  - repository/
  - task-manager/
    - queues.go
    - executors.go
    - interface.go
    - task-manager.go
  - requester/
   - interface.go
   - repo-requester.go
  - routes/
    - routes.go
  - test/
    - mocks/
      - mock_commit_usecase.go
    - unit/
        - contollers/
            - commit_test.go
  - usecase/
  - utils/
  - .gitignore
  - go.mod
  - go.sum

```

## Dependencies

The project uses the following dependencies:

- [gorilla/mux](https://github.com/gorilla/mux) - URL router and dispatcher for Go
- [testify](https://github.com/stretchr/testify) - A toolkit with common assertions and mocks
- [mockery](https://github.com/vektra/mockery) - A mock code autogenerator for Golang
- [gorm.io/gorm](https://gorm.io/) - The fantastic ORM library for Golang
- [github.com/golang/mock](https://github.com/golang/mock) - GoMock is a mocking framework for the Go programming language.

## Usage

### Running the Application

To run the application locally:

```sh
go run cmd/main.go
```

The application will start on `http://localhost:8080`.

## Running Tests

The project includes unit tests for the controller methods. To run the tests, use the following command:

```sh
go test ./...
```

## API Endpoints

Find the documentation to the API endpoints here: https://documenter.getpostman.com/view/26825676/2sA3kPpjD1
