package store

import (
	"context"

	"github.com/google/uuid"
	"github.com/sourabh-kumar2/domain-driven-design/coffeeco/internal"
)

type Store struct {
	ID              uuid.UUID
	Location        string
	ProductsForSale []coffeelover.Product
}

type Service struct {
	repo Repository
}

func (s Service) GetStoreSpecificDiscount(ctx context.Context, storeID uuid.UUID) (float32, error) {
	dis, err := s.repo.GetStoreDiscount(ctx, storeID)
	if err != nil {
		return 0, err
	}
	return float32(dis), nil
}
