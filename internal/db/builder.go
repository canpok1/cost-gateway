package db

import (
	"fmt"
	"strings"
)

type QueryBuilder struct {
	baseSQL          string
	conditionQueries []string
	conditionValues  []interface{}
}

func NewQueryBuilder(baseSQL string) *QueryBuilder {
	return &QueryBuilder{
		baseSQL:          baseSQL,
		conditionQueries: []string{},
		conditionValues:  []interface{}{},
	}
}

func (b *QueryBuilder) AddCondition(q string, v interface{}) *QueryBuilder {
	b.conditionQueries = append(b.conditionQueries, q)
	b.conditionValues = append(b.conditionValues, v)
	return b
}

func (b *QueryBuilder) Build() (sql string, values []interface{}) {
	whereQuery := strings.Join(b.conditionQueries, " AND ")
	if whereQuery == "" {
		sql = b.baseSQL
	} else {
		sql = fmt.Sprintf("%s WHERE %s", b.baseSQL, whereQuery)
	}
	values = b.conditionValues
	return
}
