package api

import (
	"context"

	"github.com/canpok1/code-gateway/internal/db"
)

func getApiV1CostsTypes(ctx context.Context, client *db.Queries) (*GetCostsTypesResponse, error) {
	records, err := client.GetCostTypes(ctx)
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
