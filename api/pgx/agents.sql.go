// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: agents.sql

package pgx

import (
	"context"
)

const createAgent = `-- name: CreateAgent :one
INSERT INTO agents (name, email)
VALUES ($1, $2)
RETURNING id, name, email
`

type CreateAgentParams struct {
	Name  string
	Email string
}

// CreateAgent
//
//	INSERT INTO agents (name, email)
//	VALUES ($1, $2)
//	RETURNING id, name, email
func (q *Queries) CreateAgent(ctx context.Context, arg CreateAgentParams) (*Agent, error) {
	row := q.db.QueryRow(ctx, createAgent, arg.Name, arg.Email)
	var i Agent
	err := row.Scan(&i.ID, &i.Name, &i.Email)
	return &i, err
}

const deleteAgent = `-- name: DeleteAgent :one
DELETE FROM agents
WHERE id = $1
RETURNING id, name, email
`

// DeleteAgent
//
//	DELETE FROM agents
//	WHERE id = $1
//	RETURNING id, name, email
func (q *Queries) DeleteAgent(ctx context.Context, id int64) (*Agent, error) {
	row := q.db.QueryRow(ctx, deleteAgent, id)
	var i Agent
	err := row.Scan(&i.ID, &i.Name, &i.Email)
	return &i, err
}

const getAgent = `-- name: GetAgent :one
SELECT id, name, email FROM agents
WHERE id = $1
`

// GetAgent
//
//	SELECT id, name, email FROM agents
//	WHERE id = $1
func (q *Queries) GetAgent(ctx context.Context, id int64) (*Agent, error) {
	row := q.db.QueryRow(ctx, getAgent, id)
	var i Agent
	err := row.Scan(&i.ID, &i.Name, &i.Email)
	return &i, err
}

const listAgents = `-- name: ListAgents :many
SELECT id, name, email FROM agents
ORDER BY name
`

// ListAgents
//
//	SELECT id, name, email FROM agents
//	ORDER BY name
func (q *Queries) ListAgents(ctx context.Context) ([]*Agent, error) {
	rows, err := q.db.Query(ctx, listAgents)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Agent
	for rows.Next() {
		var i Agent
		if err := rows.Scan(&i.ID, &i.Name, &i.Email); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listAgentsByAuthorIDs = `-- name: ListAgentsByAuthorIDs :many
SELECT agents.id, agents.name, agents.email, authors.id AS author_id FROM agents, authors
WHERE agents.id = authors.agent_id AND authors.id  = ANY($1::bigint[])
`

type ListAgentsByAuthorIDsRow struct {
	ID       int64
	Name     string
	Email    string
	AuthorID int64
}

// ListAgentsByAuthorIDs
//
//	SELECT agents.id, agents.name, agents.email, authors.id AS author_id FROM agents, authors
//	WHERE agents.id = authors.agent_id AND authors.id  = ANY($1::bigint[])
func (q *Queries) ListAgentsByAuthorIDs(ctx context.Context, dollar_1 []int64) ([]*ListAgentsByAuthorIDsRow, error) {
	rows, err := q.db.Query(ctx, listAgentsByAuthorIDs, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*ListAgentsByAuthorIDsRow
	for rows.Next() {
		var i ListAgentsByAuthorIDsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.AuthorID,
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

const updateAgent = `-- name: UpdateAgent :one
UPDATE agents
SET name = $2, email = $3
WHERE id = $1
RETURNING id, name, email
`

type UpdateAgentParams struct {
	ID    int64
	Name  string
	Email string
}

// UpdateAgent
//
//	UPDATE agents
//	SET name = $2, email = $3
//	WHERE id = $1
//	RETURNING id, name, email
func (q *Queries) UpdateAgent(ctx context.Context, arg UpdateAgentParams) (*Agent, error) {
	row := q.db.QueryRow(ctx, updateAgent, arg.ID, arg.Name, arg.Email)
	var i Agent
	err := row.Scan(&i.ID, &i.Name, &i.Email)
	return &i, err
}
