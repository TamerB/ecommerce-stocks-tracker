-- name: ListProductStocksBySKU :many
SELECT 
    products.id AS id, 
    products.sku AS sku, 
    products.name AS name, 
    products.created_at AS created_at, 
    products.updated_at AS updated_at,
    countries.country_code AS country,
    stocks.quantity AS quantity
FROM products
LEFT JOIN stocks
ON products.id = stocks.product_id
LEFT JOIN countries
ON countries.id = stocks.country_id
WHERE products.sku = $1;

-- name: GetProductStockCountBySKU :one
SELECT SUM(stocks.quantity) AS total 
FROM stocks
INNER JOIN products
ON stocks.product_id = products.id
AND products.sku = $1 LIMIT 1;

-- name: GetStockByProductSKUAndCountryCodeForUpdate :one
SELECT stocks.id, stocks.quantity 
FROM stocks
INNER JOIN products
ON stocks.product_id = products.id
AND products.sku = sqlc.arg(sku)
INNER JOIN countries
ON stocks.country_id = countries.id
AND countries.country_code = sqlc.arg(country)
LIMIT 1
FOR NO KEY UPDATE;

-- name: ConsumeStock :exec
UPDATE stocks
SET quantity = quantity - sqlc.arg(quantity)
WHERE id = sqlc.arg(id);