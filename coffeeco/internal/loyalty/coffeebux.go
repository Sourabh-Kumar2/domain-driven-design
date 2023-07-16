package loyalty

import (
	"github.com/google/uuid"
	"github.com/sourabh-kumar2/domain-driven-design/coffeeco/internal"
	"github.com/sourabh-kumar2/domain-driven-design/coffeeco/internal/store"
)

type CoffeeBux struct {
	ID                                    uuid.UUID
	store                                 store.Store
	coffeeLover                           coffeelover.CoffeeLover
	FreeDrinksAvailable                   int
	RemainingDrinkPurchasesUntilFreeDrink int
}
