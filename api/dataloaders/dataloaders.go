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
		AgentByAuthorID *AgentLoader
	}
)

const key = contextKey("dataloaders")

func newLoaders(ctx context.Context, repoSvc pgx.Repository) *Loaders {
	return &Loaders{
		AgentByAuthorID: newAgentByAuthorID(ctx, repoSvc),
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
