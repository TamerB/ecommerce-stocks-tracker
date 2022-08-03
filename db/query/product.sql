-- name: GetProductBySKU :one
SELECT * FROM products WHERE sku = $1 LIMIT 1;