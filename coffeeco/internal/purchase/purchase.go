package purchase

import (
	"time"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
	"github.com/sourabh-kumar2/domain-driven-design/coffeeco/internal"
	"github.com/sourabh-kumar2/domain-driven-design/coffeeco/internal/payment"
	"github.com/sourabh-kumar2/domain-driven-design/coffeeco/internal/store"
)

type Purchase struct {
	id                 uuid.UUID
	Store              store.Store
	ProductsToPurchase []coffeelover.Product
	total              money.Money
	PaymentMeans       payment.Means
	timeOfPurchase     time.Time
	CardToken          *string
}
