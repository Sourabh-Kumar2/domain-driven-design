package store

import (
	"github.com/google/uuid"
	coffeelover "github.com/sourabh-kumar2/domain-driven-design/coffeeco/internal"
)

type Store struct {
	ID              uuid.UUID
	Location        string
	ProductsForSale []coffeelover.Product
}
