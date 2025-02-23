package db

import (
	"context"
	"fmt"
	"log"
	"strings"
)

type FindMonthlyCostsCondition struct {
	CostTypeID uint64
	BeginYear  *uint32
	BeginMonth *uint32
	EndYear    *uint32
	EndMonth   *uint32
}

func (q *Queries) FindMonthlyCostsByCondition(ctx context.Context, cond *FindMonthlyCostsCondition) ([]FindMonthlyCostsRow, error) {
	whereConditions := []string{}
	whereValues := []interface{}{}

	whereConditions = append(whereConditions, "cost_type_id = ?")
	whereValues = append(whereValues, cond.CostTypeID)

	if cond.BeginYear != nil {
		whereConditions = append(whereConditions, "begin_year = ?")
		whereValues = append(whereValues, cond.BeginYear)
	}
	if cond.BeginMonth != nil {
		whereConditions = append(whereConditions, "begin_month = ?")
		whereValues = append(whereValues, cond.BeginMonth)
	}
	if cond.EndYear != nil {
		whereConditions = append(whereConditions, "end_year = ?")
		whereValues = append(whereValues, cond.EndYear)
	}
	if cond.EndMonth != nil {
		whereConditions = append(whereConditions, "end_month = ?")
		whereValues = append(whereValues, cond.EndMonth)
	}

	sql := fmt.Sprintf("%s WHERE %s", findMonthlyCosts, strings.Join(whereConditions, " AND "))
	log.Println(sql)

	rows, err := q.db.QueryContext(ctx, sql, whereValues...)
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
