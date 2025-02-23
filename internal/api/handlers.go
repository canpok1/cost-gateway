package api

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/canpok1/code-gateway/internal/db"
)

func (s *Server) getApiV1CostsTypes(ctx context.Context) (*GetCostsTypesResponse, error) {
	q := db.New(s.db)
	records, err := q.FindCostTypeAll(ctx)
	if err != nil {
		return nil, err
	}

	costTypes := []CostType{}
	for _, record := range records {
		costTypes = append(costTypes, CostType{
			CostTypeId:   int64(record.ID),
			CostTypeName: record.TypeName,
		})
	}

	return &GetCostsTypesResponse{
		CostTypes: &costTypes,
	}, nil
}

func (s *Server) getApiV1CostsMonthly(ctx context.Context, param *GetApiV1CostsMonthlyParams) (*GetMonthlyCostsResponse, error) {
	q := db.New(s.db)
	cond := db.FindMonthlyCostsCondition{CostTypeID: uint64(param.CostTypeId)}
	if param.BeginYear != nil {
		v := uint32(*param.BeginYear)
		cond.BeginYear = &v
	}
	if param.BeginMonth != nil {
		v := uint32(*param.BeginMonth)
		cond.BeginMonth = &v
	}
	if param.EndYear != nil {
		v := uint32(*param.EndYear)
		cond.EndYear = &v
	}
	if param.EndMonth != nil {
		v := uint32(*param.EndMonth)
		cond.EndMonth = &v
	}

	records, err := q.FindMonthlyCostsByCondition(ctx, &cond)
	if err != nil {
		return nil, fmt.Errorf("failed to find monthly costs: %w", err)
	}

	resp := GetMonthlyCostsResponse{
		Costs: []Cost{},
	}
	for _, r := range records {
		resp.Costs = append(resp.Costs, Cost{
			CostTypeId:   int64(r.CostTypeID),
			CostTypeName: r.TypeName,
			Month:        int32(r.CostMonth),
			Year:         int32(r.CostYear),
			Yen:          int32(r.CostYen),
		})
	}

	return &resp, nil
}

func (s *Server) postApiV1CostsMonthly(ctx context.Context, body *PostApiV1CostsMonthlyJSONRequestBody) (*PostMonthlyCostsResponse, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	q := db.New(s.db).WithTx(tx)

	costType, err := q.FindCostTypeByTypeName(ctx, body.CostTypeName)
	if err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("failed to find types: %w", err)
	}

	var costTypeID int64
	if err != sql.ErrNoRows {
		costTypeID = int64(costType.ID)
	} else {
		result, err := q.InsertCostType(ctx, body.CostTypeName)
		if err != nil {
			return nil, fmt.Errorf("failed to insert costType: %w", err)
		}

		costTypeID, err = result.LastInsertId()
		if err != nil {
			return nil, fmt.Errorf("failed to get last insert id: %w", err)
		}
	}

	for _, cost := range body.Costs {
		_, err := q.UpsertMonthlyCost(ctx, db.UpsertMonthlyCostParams{
			CostTypeID: uint64(costTypeID),
			CostYear:   uint32(*cost.Year),
			CostMonth:  uint32(*cost.Month),
			CostYen:    uint32(*cost.Yen),
		})
		if err != nil {
			return nil, fmt.Errorf("failed to insert monthly cost: %w", err)
		}
	}

	resp := PostMonthlyCostsResponse{
		CostTypeId: costTypeID,
		Count:      int32(len(body.Costs)),
	}

	return &resp, tx.Commit()
}
