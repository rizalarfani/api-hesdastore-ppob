package repositories

import (
	"context"
	"hesdastore/api-ppob/domain/model"

	"github.com/jmoiron/sqlx"
)

type BillingRepository interface {
	CreateInquiry(ctx context.Context, tx *sqlx.Tx, order *model.InquiryBilling) (*model.InquiryBilling, error)
}
