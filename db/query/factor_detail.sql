-- name: GetFactorDetail :one
SELECT * FROM factor_details
WHERE id = ? LIMIT 1;

-- name: ListFactorDetails :many
SELECT * FROM factor_details
WHERE factor_id = ?
ORDER BY id
LIMIT ?
OFFSET ?;

-- name: CreateFactorDetail :execresult
INSERT INTO factor_details (
    factor_id, product_id, sale_count, sale_price
) VALUES (
             ?, ?, ?, ?
         );

-- name: UpdateFactorDetail :execresult
UPDATE factor_details
SET product_id=?, sale_count=?, sale_price=?
WHERE id = ?;

-- name: DeleteFactorDetail :exec
DELETE FROM factor_details
WHERE id = ?;