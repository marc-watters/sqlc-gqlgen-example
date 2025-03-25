// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: authors.sql

package pgx

import (
	"context"
)

const createAuthor = `-- name: CreateAuthor :one
INSERT INTO authors (name, website, agent_id)
VALUES ($1, $2, $3)
RETURNING id, name, website, agent_id
`

type CreateAuthorParams struct {
	Name    string
	Website *string
	AgentID int64
}

// CreateAuthor
//
//	INSERT INTO authors (name, website, agent_id)
//	VALUES ($1, $2, $3)
//	RETURNING id, name, website, agent_id
func (q *Queries) CreateAuthor(ctx context.Context, arg CreateAuthorParams) (*Author, error) {
	row := q.db.QueryRow(ctx, createAuthor, arg.Name, arg.Website, arg.AgentID)
	var i Author
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Website,
		&i.AgentID,
	)
	return &i, err
}

const deleteAuthor = `-- name: DeleteAuthor :one
DELETE FROM authors
WHERE id = $1
RETURNING id, name, website, agent_id
`

// DeleteAuthor
//
//	DELETE FROM authors
//	WHERE id = $1
//	RETURNING id, name, website, agent_id
func (q *Queries) DeleteAuthor(ctx context.Context, id int64) (*Author, error) {
	row := q.db.QueryRow(ctx, deleteAuthor, id)
	var i Author
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Website,
		&i.AgentID,
	)
	return &i, err
}

const getAuthor = `-- name: GetAuthor :one
SELECT id, name, website, agent_id FROM authors
WHERE id = $1
`

// GetAuthor
//
//	SELECT id, name, website, agent_id FROM authors
//	WHERE id = $1
func (q *Queries) GetAuthor(ctx context.Context, id int64) (*Author, error) {
	row := q.db.QueryRow(ctx, getAuthor, id)
	var i Author
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Website,
		&i.AgentID,
	)
	return &i, err
}

const listAuthors = `-- name: ListAuthors :many
SELECT id, name, website, agent_id FROM authors
ORDER BY name
`

// ListAuthors
//
//	SELECT id, name, website, agent_id FROM authors
//	ORDER BY name
func (q *Queries) ListAuthors(ctx context.Context) ([]*Author, error) {
	rows, err := q.db.Query(ctx, listAuthors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Website,
			&i.AgentID,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listAuthorsByAgentID = `-- name: ListAuthorsByAgentID :many
SELECT authors.id, authors.name, authors.website, authors.agent_id FROM authors, agents
WHERE agents.id = authors.agent_id AND authors.agent_id = $1
`

// ListAuthorsByAgentID
//
//	SELECT authors.id, authors.name, authors.website, authors.agent_id FROM authors, agents
//	WHERE agents.id = authors.agent_id AND authors.agent_id = $1
func (q *Queries) ListAuthorsByAgentID(ctx context.Context, agentID int64) ([]*Author, error) {
	rows, err := q.db.Query(ctx, listAuthorsByAgentID, agentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Website,
			&i.AgentID,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listAuthorsByAgentIDs = `-- name: ListAuthorsByAgentIDs :many
SELECT authors.id, authors.name, authors.website, authors.agent_id FROM authors, agents
WHERE authors.agent_id = agents.id AND agents.id = ANY($1::bigint[])
`

// ListAuthorsByAgentIDs
//
//	SELECT authors.id, authors.name, authors.website, authors.agent_id FROM authors, agents
//	WHERE authors.agent_id = agents.id AND agents.id = ANY($1::bigint[])
func (q *Queries) ListAuthorsByAgentIDs(ctx context.Context, dollar_1 []int64) ([]*Author, error) {
	rows, err := q.db.Query(ctx, listAuthorsByAgentIDs, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Website,
			&i.AgentID,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listAuthorsByBookID = `-- name: ListAuthorsByBookID :many
SELECT authors.id, authors.name, authors.website, authors.agent_id FROM authors, book_authors
WHERE authors.id = book_authors.author_id AND book_authors.book_id = $1
`

// ListAuthorsByBookID
//
//	SELECT authors.id, authors.name, authors.website, authors.agent_id FROM authors, book_authors
//	WHERE authors.id = book_authors.author_id AND book_authors.book_id = $1
func (q *Queries) ListAuthorsByBookID(ctx context.Context, bookID int64) ([]*Author, error) {
	rows, err := q.db.Query(ctx, listAuthorsByBookID, bookID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Website,
			&i.AgentID,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listAuthorsByBookIDs = `-- name: ListAuthorsByBookIDs :many
SELECT authors.id, authors.name, authors.website, authors.agent_id, book_authors.book_id FROM authors, book_authors
WHERE book_authors.author_id = authors.id AND book_authors.book_id = ANY($1::bigint[])
`

type ListAuthorsByBookIDsRow struct {
	ID      int64
	Name    string
	Website *string
	AgentID int64
	BookID  int64
}

// ListAuthorsByBookIDs
//
//	SELECT authors.id, authors.name, authors.website, authors.agent_id, book_authors.book_id FROM authors, book_authors
//	WHERE book_authors.author_id = authors.id AND book_authors.book_id = ANY($1::bigint[])
func (q *Queries) ListAuthorsByBookIDs(ctx context.Context, dollar_1 []int64) ([]*ListAuthorsByBookIDsRow, error) {
	rows, err := q.db.Query(ctx, listAuthorsByBookIDs, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*ListAuthorsByBookIDsRow
	for rows.Next() {
		var i ListAuthorsByBookIDsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Website,
			&i.AgentID,
			&i.BookID,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAuthor = `-- name: UpdateAuthor :one
UPDATE authors
SET name = $2, website = $3, agent_id = $4
WHERE id = $1
RETURNING id, name, website, agent_id
`

type UpdateAuthorParams struct {
	ID      int64
	Name    string
	Website *string
	AgentID int64
}

// UpdateAuthor
//
//	UPDATE authors
//	SET name = $2, website = $3, agent_id = $4
//	WHERE id = $1
//	RETURNING id, name, website, agent_id
func (q *Queries) UpdateAuthor(ctx context.Context, arg UpdateAuthorParams) (*Author, error) {
	row := q.db.QueryRow(ctx, updateAuthor,
		arg.ID,
		arg.Name,
		arg.Website,
		arg.AgentID,
	)
	var i Author
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Website,
		&i.AgentID,
	)
	return &i, err
}
