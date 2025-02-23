-- name: UpsertMonthlyCost :execresult
INSERT INTO monthly_costs (
    cost_type_id, cost_year, cost_month, cost_yen
) VALUES (?, ?, ?, ?)
ON DUPLICATE KEY UPDATE
    cost_yen = VALUES(cost_yen)
;
