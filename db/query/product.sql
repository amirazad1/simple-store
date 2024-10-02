-- name: GetProduct :one
SELECT * FROM products
WHERE id = ? LIMIT 1
FOR UPDATE;

-- name: ListProducts :many
SELECT * FROM products
ORDER BY name
LIMIT ?
OFFSET ?;

-- name: CreateProduct :execresult
INSERT INTO products (
    name, brand, model, init_number, present_number, buy_price, buy_date, sale_price
) VALUES (
             ?, ?, ?, ?, ?, ?, ?, ?
         );

-- name: UpdateProduct :execresult
UPDATE products
SET name=?,brand=?,model=?,init_number=?,present_number=?,buy_price=?,buy_date=?,sale_price=?
WHERE id = ?;

-- name: UpdateProductPresent :execresult
UPDATE products
SET present_number=?
WHERE id = ?;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = ?;