package services

import (
	"context"
	"hesdastore/api-ppob/domain/dto"
	"hesdastore/api-ppob/domain/model"
)

type BillingService interface {
	Inquiry(context.Context, *dto.InquiryBillRequest, *model.ApiUser) (*dto.InquiryBillingResponse, error)
}
