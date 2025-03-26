-- name: GetAgent :one
SELECT * FROM agents
WHERE id = $1;

-- name: ListAgents :many
SELECT * FROM agents
ORDER BY name;

-- name: ListAgentsByAuthorIDs :many
SELECT agents.*, authors.id AS author_id FROM agents, authors
WHERE agents.id = authors.agent_id AND authors.id  = ANY($1::bigint[]);

-- name: CreateAgent :one
INSERT INTO agents (name, email)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateAgent :one
UPDATE agents
SET name = $2, email = $3
WHERE id = $1
RETURNING *;

-- name: DeleteAgent :one
DELETE FROM agents
WHERE id = $1
RETURNING *;
