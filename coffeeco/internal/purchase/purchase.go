package purchase

import (
	"context"
	"errors"
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
	cardToken          *string
}

func (p *Purchase) validateAndEnrich() error {
	if len(p.ProductsToPurchase) == 0 {
		return errors.New("purchase must consist of at least one product")
	}
	p.total = *money.New(0, "USD")
	for _, v := range p.ProductsToPurchase {
		newTotal, _ := p.total.Add(&v.BasePrice)
		p.total = *newTotal
	}
	if p.total.IsZero() {
		return errors.New("likely mistake; purchase should never be 0. Please validate")
	}
	p.id = uuid.New()
	p.timeOfPurchase = time.Now()
	return nil
}

type CardChargeService interface {
	ChargeCard(ctx context.Context, amount money.Money, cardToken string) error
}

type CashChargeService interface {
	CollectCash(ctx context.Context, amount money.Money) error
}

type Service struct {
	cardService  CardChargeService
	cashService  CashChargeService
	purchaseRepo Repository
}

func (s Service) CompletePurchase(ctx context.Context, purchase *Purchase) error {
	if err := purchase.validateAndEnrich(); err != nil {
		return err
	}
	switch purchase.PaymentMeans {
	case payment.MEANS_CARD:
		if err := s.cardService.ChargeCard(ctx, purchase.total, *purchase.cardToken); err != nil {
			return errors.New("card charge failed, cancelling purchase")
		}
	case payment.MEANS_CASH:
		if err := s.cashService.CollectCash(ctx, purchase.total); err != nil {
			return errors.New("card charge failed, cancelling purchase")
		}
	default:
		return errors.New("unknown payment type")
	}
	if err := s.purchaseRepo.Store(ctx, *purchase); err != nil {
		return errors.New("failed to store purchase")
	}
	return nil
}
