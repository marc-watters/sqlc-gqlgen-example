-- name: SetBookAuthor :exec
INSERT INTO book_authors (book_id, author_id)
VALUES ($1, $2);

-- name: UnsetBookAuthors :exec
DELETE FROM book_authors
WHERE book_id = $1;
