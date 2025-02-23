package db

import (
	"context"
	"log"
)

type FindMonthlyCostsCondition struct {
	CostTypeID uint64
	BeginYear  *uint32
	BeginMonth *uint32
	EndYear    *uint32
	EndMonth   *uint32
}

func (q *Queries) FindMonthlyCostsByCondition(ctx context.Context, cond *FindMonthlyCostsCondition) ([]FindMonthlyCostsRow, error) {
	builder := NewQueryBuilder(findMonthlyCosts)

	builder.AddCondition("cost_type_id = ?", cond.CostTypeID)

	if cond.BeginYear != nil {
		builder.AddCondition("cost_year >= ?", cond.BeginYear)
	}
	if cond.BeginMonth != nil {
		builder.AddCondition("cost_month >= ?", cond.BeginMonth)
	}
	if cond.EndYear != nil {
		builder.AddCondition("cost_year <= ?", cond.EndYear)
	}
	if cond.EndMonth != nil {
		builder.AddCondition("cost_month <= ?", cond.EndMonth)
	}

	sql, values := builder.Build()
	log.Println(sql)

	rows, err := q.db.QueryContext(ctx, sql, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindMonthlyCostsRow
	for rows.Next() {
		var i FindMonthlyCostsRow
		if err := rows.Scan(
			&i.CostTypeID,
			&i.CostYear,
			&i.CostMonth,
			&i.CostYen,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.TypeName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
