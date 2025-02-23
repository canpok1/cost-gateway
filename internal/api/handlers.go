package api

import (
	"context"
)

func (s *Server) getApiV1CostsTypes(ctx context.Context) (*GetCostsTypesResponse, error) {
	records, err := s.client.GetCostTypes(ctx)
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
