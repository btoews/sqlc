-- name: ListFoo :many
SELECT id, bar
FROM foo
WHERE id = $1 AND bar = $2;

-- name: ListBaz :many
SELECT * FROM foo;

