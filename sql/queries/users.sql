--name : CreateUser :one
-- here $1, $2 etc.. with take in the arguments values which is 
-- the syntax for sqlc
INSERT INTO users (id, created_at, updated_at, name)
VALUES ($1, $2, $3, $4)
RETURNING *;