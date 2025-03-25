package gqlgen

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.68

import (
	"context"
	"fmt"

	"github.com/marc-watters/sqlc-gqlgen-example/v2/pgx"
)

// Authors is the resolver for the authors field.
func (r *agentResolver) Authors(ctx context.Context, obj *pgx.Agent) ([]*pgx.Author, error) {
	panic(fmt.Errorf("not implemented: Authors - authors"))
}

// Agent returns AgentResolver implementation.
func (r *Resolver) Agent() AgentResolver { return &agentResolver{r} }

type agentResolver struct{ *Resolver }
