/* name: CreateOrder :exec */
INSERT INTO orders (id, amount, status, user_id, created_at)
VALUES (?, ?, ?, ?, ?);

/* name: CreateOrderDetails :exec */
INSERT INTO order_details (order_id, product_id, quantity)
VALUES (?, ?, ?);

/* name: GetOrdersPaginated :many */
SELECT id, amount, status, user_id, created_at
FROM orders
ORDER BY
    CASE WHEN sqlc.arg(sort_by) = 'createdAt' AND sqlc.arg(sort_order) = 'ASC' THEN created_at END ,
    CASE WHEN sqlc.arg(sort_by) = 'createdAt' AND sqlc.arg(sort_order) = 'DESC' THEN created_at END DESC ,
    CASE WHEN sqlc.arg(sort_by) = 'status' AND sqlc.arg(sort_order) = 'ASC' THEN status END,
    CASE WHEN sqlc.arg(sort_by) = 'status' AND sqlc.arg(sort_order) = 'DESC' THEN status END DESC,
    CASE WHEN sqlc.arg(sort_by) = 'amount' AND sqlc.arg(sort_order) = 'ASC' THEN amount END,
    CASE WHEN sqlc.arg(sort_by) = 'amount' AND sqlc.arg(sort_order) = 'DESC' THEN amount END DESC
LIMIT ?
OFFSET ?;

/* name: GetOrderTotalCount :one */
SELECT COUNT(*)
FROM orders
WHERE STRCMP(sqlc.arg(user_id),user_id) = 0;

/* name: GetOrderById :one */
SELECT id, amount, status, user_id, created_at
FROM orders
WHERE STRCMP(sqlc.arg(user_id),user_id) = 0 AND STRCMP(sqlc.arg(id),id) = 0;

/* name: GetOrderDetails :many */
SELECT order_id, product_id, quantity
FROM order_details
WHERE order_id = ?;