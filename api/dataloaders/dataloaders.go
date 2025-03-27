package dataloaders

//go:generate dataloaden AgentLoader int64 *github.com/marc-watters/sqlc-gqlgen-example/v2/pgx.Agent
//go:generate dataloaden AuthorSliceLoader int64 []*github.com/marc-watters/sqlc-gqlgen-example/v2/pgx.Author

import (
	"context"

	"github.com/marc-watters/sqlc-gqlgen-example/v2/pgx"
)

type (
	contextKey string

	Loaders struct{}
)

const key = contextKey("dataloaders")

func newLoaders(ctx context.Context, repoSvc pgx.Repository) *Loaders {
	return &Loaders{}
}
