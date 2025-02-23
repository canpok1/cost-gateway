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
