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

-- name: FindMonthlyCosts :many
SELECT mc.*, ct.type_name
FROM monthly_costs mc
JOIN cost_types ct ON mc.cost_type_id = ct.id
;
