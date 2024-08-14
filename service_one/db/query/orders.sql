-- -- name: CreateOrder : one
-- INSERT INTO orders (
--     user_id,
--     item_id,
--     status,
--     payment_link
-- )   VALUES (
--     $1, $2, $3, $4
-- ) RETURNING *;

-- -- name: GetOrder :one
-- SELECT * FROM orders
-- WHERE order_id =  $1;

-- -- name: UpdateOrder :one
-- UPDATE orders
-- SET 
--     status = COLAESCE(sqlc.narg(status), status),
--     payment_link = COLAESCE(sqlc.narg(payment_link), payment_link),
-- WHERE
--     order_id = sqlc.arg(order_id)    
-- RETURNING *;

-- -- name: DeleteOrder :exec
-- DELETE FROM orders
-- WHERE order_id = $1;
