package dataloaders

//go:generate dataloaden AgentLoader int64 *github.com/marc-watters/sqlc-gqlgen-example/v2/pgx.Agent
//go:generate dataloaden AuthorSliceLoader int64 []*github.com/marc-watters/sqlc-gqlgen-example/v2/pgx.Author
//go:generate dataloaden BookSliceLoader int64 []*github.com/marc-watters/sqlc-gqlgen-example/v2/pgx.Book

import (
	"context"
	"time"

	"github.com/marc-watters/sqlc-gqlgen-example/v2/pgx"
)

type (
	contextKey string

	Loaders struct {
		AgentByAuthorID  *AgentLoader
		AuthorsByAgentID *AuthorSliceLoader
		AuthorsByBookID  *AuthorSliceLoader
	}
)

const key = contextKey("dataloaders")

func newLoaders(ctx context.Context, repoSvc pgx.Repository) *Loaders {
	return &Loaders{
		AgentByAuthorID:  newAgentByAuthorID(ctx, repoSvc),
		AuthorsByAgentID: newAuthorsByAgentID(ctx, repoSvc),
		AuthorsByBookID:  newAuthorsByBookID(ctx, repoSvc),
	}
}

func newAgentByAuthorID(ctx context.Context, repo pgx.Repository) *AgentLoader {
	return NewAgentLoader(AgentLoaderConfig{
		MaxBatch: 100,
		Wait:     5 * time.Millisecond,
		Fetch: func(authorIDs []int64) ([]*pgx.Agent, []error) {
			// db query
			res, err := repo.ListAgentsByAuthorIDs(ctx, authorIDs)
			if err != nil {
				return nil, []error{err}
			}
			// map
			groupByAuthorID := make(map[int64]*pgx.Agent, len(authorIDs))
			for _, r := range res {
				groupByAuthorID[r.AuthorID] = &pgx.Agent{
					ID:    r.ID,
					Name:  r.Name,
					Email: r.Email,
				}
			}
			// order
			result := make([]*pgx.Agent, len(authorIDs))
			for i, authorID := range authorIDs {
				result[i] = groupByAuthorID[authorID]
			}
			return result, nil
		},
	})
}

func newAuthorsByAgentID(ctx context.Context, repo pgx.Repository) *AuthorSliceLoader {
	return NewAuthorSliceLoader(AuthorSliceLoaderConfig{
		MaxBatch: 100,
		Wait:     5 * time.Millisecond,
		Fetch: func(agentIDs []int64) ([][]*pgx.Author, []error) {
			// db query
			res, err := repo.ListAuthorsByAgentIDs(ctx, agentIDs)
			if err != nil {
				return nil, []error{err}
			}
			// group
			groupByAgentID := make(map[int64][]*pgx.Author, len(agentIDs))
			for _, r := range res {
				groupByAgentID[r.AgentID] = append(groupByAgentID[r.AgentID], r)
			}
			// order
			result := make([][]*pgx.Author, len(agentIDs))
			for i, agentID := range agentIDs {
				result[i] = groupByAgentID[agentID]
			}
			return result, nil
		},
	})
}

func newAuthorsByBookID(ctx context.Context, repo pgx.Repository) *AuthorSliceLoader {
	return NewAuthorSliceLoader(AuthorSliceLoaderConfig{
		MaxBatch: 100,
		Wait:     5 * time.Millisecond,
		Fetch: func(bookIDs []int64) ([][]*pgx.Author, []error) {
			// db query
			res, err := repo.ListAuthorsByBookIDs(ctx, bookIDs)
			if err != nil {
				return nil, []error{err}
			}
			// group
			groupByBookID := make(map[int64][]*pgx.Author, len(bookIDs))
			for _, r := range res {
				groupByBookID[r.BookID] = append(groupByBookID[r.BookID], &pgx.Author{
					ID:      r.ID,
					Name:    r.Name,
					Website: r.Website,
					AgentID: r.AgentID,
				})
			}
			// order
			result := make([][]*pgx.Author, len(bookIDs))
			for i, bookID := range bookIDs {
				result[i] = groupByBookID[bookID]
			}
			return result, nil
		},
	})
}
