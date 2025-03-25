package gqlgen

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.68

import (
	"context"

	"github.com/marc-watters/sqlc-gqlgen-example/v2/pgx"
)

// Agent is the resolver for the agent field.
func (r *queryResolver) Agent(ctx context.Context, id int64) (*pgx.Agent, error) {
	return r.Repository.GetAgent(ctx, id)
}

// Agents is the resolver for the agents field.
func (r *queryResolver) Agents(ctx context.Context) ([]*pgx.Agent, error) {
	return r.Repository.ListAgents(ctx)
}

// Author is the resolver for the author field.
func (r *queryResolver) Author(ctx context.Context, id int64) (*pgx.Author, error) {
	return r.Repository.GetAuthor(ctx, id)
}

// Authors is the resolver for the authors field.
func (r *queryResolver) Authors(ctx context.Context) ([]*pgx.Author, error) {
	return r.Repository.ListAuthors(ctx)
}

// Book is the resolver for the book field.
func (r *queryResolver) Book(ctx context.Context, id int64) (*pgx.Book, error) {
	return r.Repository.GetBook(ctx, id)
}

// Books is the resolver for the books field.
func (r *queryResolver) Books(ctx context.Context) ([]*pgx.Book, error) {
	return r.Repository.ListBooks(ctx)
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
