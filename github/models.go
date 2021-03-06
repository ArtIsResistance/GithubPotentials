package github

import "github.com/google/go-github/github"
import "time"

type Repository struct {
	Owner       string
	Name        string
	Description string
	Homepage    string
	License     string
	Language    string
	Commits     int
	Stars       int
	Contribs    int
}

type Commit struct {
	CommitterID int
}

type Stargazer struct {
	StarredAt time.Time
}

func castRepository(src github.Repository) Repository {
	result := Repository{
		Owner: *src.Owner.Login,
		Name:  *src.Name,
	}
	if src.Description != nil {
		result.Description = *src.Description
	}
	if src.Homepage != nil {
		result.Homepage = *src.Homepage
	}
	if src.Language != nil {
		result.Language = *src.Language
	}
	if src.License != nil && src.License.Name != nil {
		result.License = *src.License.Name
	}

	return result
}
