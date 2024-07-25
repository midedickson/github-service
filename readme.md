# GitHub Service

## Overview

GitHub Service is a Go-based application that interacts with the GitHub API to manage repositories and their commits. The service allows users to register their GitHub usernames, fetch repository information, and retrieve commit details for repositories.

## Table of Contents

1. [Getting Started](#getting-started)
2. [Dependencies](#dependencies)
3. [API Endpoints](#api-endpoints)
4. [Usage](#usage)
5. [Video Explanation](#video-example)
6. [Running Tests](#running-tests)

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

4. Create new env file:
   ```sh
   cp .env.sample .env
   ```

### Running the Application

To run the application locally:

```sh
go run cmd/main.go
```

The application will start on `http://localhost:8080`.

## Dependencies

The project uses the following dependencies:

- [gorilla/mux](https://github.com/gorilla/mux) - URL router and dispatcher for Go
- [testify](https://github.com/stretchr/testify) - A toolkit with common assertions and mocks
- [mockery](https://github.com/vektra/mockery) - A mock code autogenerator for Golang
- [gorm.io/gorm](https://gorm.io/) - The fantastic ORM library for Golang
- [github.com/golang/mock](https://github.com/golang/mock) - GoMock is a mocking framework for the Go programming language.

## API Endpoints

Find the documentation to the API endpoints here: https://documenter.getpostman.com/view/26825676/2sA3kPpjD1

## Usage

#### Register a New User:

Use the /register endpoint to register a new user.
Example: Register a user with username "chromium" and full name "Test User".

#### Get Repositories for a User:

Use the /{owner}/repos endpoint to fetch all repositories for a user.
Example: Fetch all repositories for the user "chromium".

#### Get Information for a Specific Repository:

Use the /{owner}/repos/{repo} endpoint to get detailed information about a specific repository.
Example: Get information for the repository "chromium" owned by "chromium".

#### Get All Commits for a Repository:

Use the /{owner}/repos/{repo}/commits endpoint to get all commits for a repository.
Example: Fetch all commits for the repository "chromium" owned by "chromium".

#### Request a Repository Reset:

Use the /{owner}/repos/{repo}/commits/reset/{reset_sha} endpoint to request a reset to a specific commit SHA.
Example: Reset the repository "chromium" owned by "chromium" to commit SHA "abcdef123456".

#### Get Top N Authors by Commits:

Use the /authors/top/{top_n} endpoint to fetch the top N authors by commit count.
Example: Get the top 3 authors by commit count.

## Video Explanation

### Folder Structure Walkthrough:

https://www.loom.com/share/67491e4aeec64cb6be3e86ef0e376176?sid=7cb07326-428f-4fc2-a9ac-f29f3b13601e

### Endpoint and Unit Test Demo:

https://www.loom.com/share/12f4fbce610a40048eab4d26d43a4bc8?sid=858c5c8d-68f4-4e08-98b0-b7ea42c370e6

## Running Tests

The project includes unit tests for the controller methods. To run the tests, use the following command:

```sh
go test ./...
```
