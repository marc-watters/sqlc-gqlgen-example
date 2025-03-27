-- name: GetAuthor :one
SELECT * FROM authors
WHERE id = $1;

-- name: ListAuthors :many
SELECT * FROM authors
ORDER BY name;

-- name: ListAuthorsByAgentID :many
SELECT authors.* FROM authors, agents
WHERE agents.id = authors.agent_id AND authors.agent_id = $1;

-- name: ListAuthorsByAgentIDs :many
SELECT authors.* FROM authors, agents
WHERE authors.agent_id = agents.id AND agents.id = ANY($1::bigint[]);

-- name: ListAuthorsByBookID :many
SELECT authors.* FROM authors, book_authors
WHERE authors.id = book_authors.author_id AND book_authors.book_id = $1;

-- name: ListAuthorsByBookIDs :many
SELECT authors.*, book_authors.book_id FROM authors, book_authors
WHERE book_authors.author_id = authors.id AND book_authors.book_id = ANY($1::bigint[]);

-- name: CreateAuthor :one
INSERT INTO authors (name, website, agent_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateAuthor :one
UPDATE authors
SET name = $2, website = $3, agent_id = $4
WHERE id = $1
RETURNING *;

-- name: DeleteAuthor :one
DELETE FROM authors
WHERE id = $1
RETURNING *;
