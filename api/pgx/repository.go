package pgx

import (
	"context"
	"fmt"

	uuid "github.com/jackc/pgx-gofrs-uuid"
	"github.com/jackc/pgx/v5"
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

// Open opens a database specified by the data source name.
// Format: host=foo port=5432 user=bar password=baz dbname=qux sslmode=disable"
func Open(ctx context.Context, uri string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(uri)
	if err != nil {
		return nil, fmt.Errorf("error opening connection pool: %w", err)
	}

	// UUID support: https://github.com/jackc/pgx/wiki/UUID-Support
	config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		uuid.Register(conn.TypeMap())
		return nil
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("error opening connection pool: %w", err)
	}

	return pool, nil
}

func (r *repoSvc) CreateBook(ctx context.Context, bookArg CreateBookParams, authorIDs []int64) (*Book, error) {
	book := new(Book)
	err := r.withTx(ctx, func(q *Queries) error {
		res, err := q.CreateBook(ctx, bookArg)
		if err != nil {
			return err
		}
		for _, authorID := range authorIDs {
			if err := q.SetBookAuthor(ctx, SetBookAuthorParams{
				BookID:   res.ID,
				AuthorID: authorID,
			}); err != nil {
				return err
			}
		}
		book = res
		return nil
	})
	return book, err
}

func (r *repoSvc) UpdateBook(ctx context.Context, bookArg UpdateBookParams, authorIDs []int64) (*Book, error) {
	book := new(Book)
	err := r.withTx(ctx, func(q *Queries) error {
		res, err := q.UpdateBook(ctx, bookArg)
		if err != nil {
			return err
		}
		if err = q.UnsetBookAuthors(ctx, res.ID); err != nil {
			return err
		}
		for _, authorID := range authorIDs {
			if err := q.SetBookAuthor(ctx, SetBookAuthorParams{
				BookID:   res.ID,
				AuthorID: authorID,
			}); err != nil {
				return err
			}
		}
		book = res
		return nil
	})
	return book, err
}

func (r *repoSvc) withTx(ctx context.Context, txFn func(*Queries) error) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}

	q := New(tx)
	err = txFn(q)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			err = fmt.Errorf("tx failed: %v, unable to rollback: %v", err, rbErr)
		}
	} else {
		err = tx.Commit(ctx)
	}

	return err
}
