-- name: GetFactor :one
SELECT * FROM factors
WHERE id = ? LIMIT 1;

-- name: ListFactors :many
SELECT * FROM factors
ORDER BY id DESC
LIMIT ?
OFFSET ?;

-- name: CreateFactor :execresult
INSERT INTO factors (
    customer_name, customer_mobile, seller
) VALUES (
             ?, ?, ?
         );

-- name: UpdateFactor :execresult
UPDATE factors
SET customer_name=?, customer_mobile=?, seller=?
WHERE id = ?;

-- name: DeleteFactor :exec
DELETE FROM factors
WHERE id = ?;