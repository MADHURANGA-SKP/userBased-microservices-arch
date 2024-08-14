-- name: CreateUser :one
INSERT INTO users (
    first_name,
    last_name,
    user_name,
    email,
    role,
    password
)   VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE user_id =  $1;

-- name: UpdateUser :one
UPDATE users
SET 
    first_name = COLAESCE(sqlc.narg(first_name), first_name),
    last_name = COLAESCE(sqlc.narg(last_name), last_name),
    user_name = COLAESCE(sqlc.narg(user_name), user_name),
    email = COLAESCE(sqlc.narg(email), email),
    is_email_verified = COALESCE(sqlc.narg(is_email_verified),is_email_verified)
WHERE
    user_id = sqlc.arg(user_id)    
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE user_id = $1;
