-- name: FindCostTypeAll :many
SELECT *
FROM cost_types
ORDER BY created_at
;

-- name: FindCostTypeByTypeName :one
SELECT *
FROM cost_types
WHERE type_name = ?
;
