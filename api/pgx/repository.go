package pgx

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type (
	// Repository is the application's data layer functionality.
	Repository interface {
		// agent queries
		CreateAgent(ctx context.Context, arg CreateAgentParams) (*Agent, error)
		DeleteAgent(ctx context.Context, id int64) (*Agent, error)
		GetAgent(ctx context.Context, id int64) (*Agent, error)
		ListAgents(ctx context.Context) ([]*Agent, error)
		ListAgentsByAuthorIDs(ctx context.Context, authorIDs []int64) ([]*ListAgentsByAuthorIDsRow, error)
		UpdateAgent(ctx context.Context, arg UpdateAgentParams) (*Agent, error)
		// author queries
		CreateAuthor(ctx context.Context, arg CreateAuthorParams) (*Author, error)
		DeleteAuthor(ctx context.Context, id int64) (*Author, error)
		GetAuthor(ctx context.Context, id int64) (*Author, error)
		ListAuthors(ctx context.Context) ([]*Author, error)
		ListAuthorsByAgentID(ctx context.Context, agentID int64) ([]*Author, error)
		ListAuthorsByAgentIDs(ctx context.Context, agentIDs []int64) ([]*Author, error)
		ListAuthorsByBookID(ctx context.Context, bookID int64) ([]*Author, error)
		ListAuthorsByBookIDs(ctx context.Context, bookIDs []int64) ([]*ListAuthorsByBookIDsRow, error)
		UpdateAuthor(ctx context.Context, arg UpdateAuthorParams) (*Author, error)
		// book queries
		CreateBook(ctx context.Context, bookArg CreateBookParams, authorIDs []int64) (*Book, error)
		UpdateBook(ctx context.Context, bookArg UpdateBookParams, authorIDs []int64) (*Book, error)
		DeleteBook(ctx context.Context, id int64) (*Book, error)
		GetBook(ctx context.Context, id int64) (*Book, error)
		ListBooks(ctx context.Context) ([]*Book, error)
		ListBooksByAuthorID(ctx context.Context, authorID int64) ([]*Book, error)
		ListBooksByAuthorIDs(ctx context.Context, authorIDs []int64) ([]*ListBooksByAuthorIDsRow, error)
	}

	repoSvc struct {
		*Queries
		db *pgxpool.Pool
	}
)

func NewRepository(db *pgxpool.Pool) Repository {
	return &repoSvc{
		Queries: New(db),
		db:      db,
	}
}
