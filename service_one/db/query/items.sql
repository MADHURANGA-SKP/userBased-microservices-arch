-- -- name: CreateOrder : one
-- INSERT INTO items (
--     user_id,
--     item_name,
--     quantity
--     price
-- )   VALUES (
--     $1, $2, $3, $4
-- ) RETURNING *;

-- -- name: GetOrder :one
-- SELECT * FROM items
-- WHERE item_id =  $1;

-- -- name: UpdateOrder :one
-- UPDATE items
-- SET 
--     item_name = COLAESCE(sqlc.narg(item_name), item_name),
--     quantity = COLAESCE(sqlc.narg(quantity), quantity),
--     price = COLAESCE(sqlc.narg(price), price)
-- WHERE
--     item_id = sqlc.arg(item_id)    
-- RETURNING *;

-- -- name: DeleteOrder :exec
-- DELETE FROM items
-- WHERE item_id = $1;
