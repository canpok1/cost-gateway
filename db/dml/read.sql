-- name: GetCostTypes :many
SELECT *
FROM cost_types
ORDER BY created_at
;
