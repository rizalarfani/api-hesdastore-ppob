package services

import (
	"context"
	"hesdastore/api-ppob/domain/dto"
	"hesdastore/api-ppob/domain/model"
)

type BillingService interface {
	Inquiry(context.Context, *dto.InquiryBillRequest, *model.ApiUser) (*dto.InquiryBillingResponse, error)
	PayBill(context.Context, *dto.PayBillRequest, *model.ApiUser) (*dto.PayBillResponse, error)
}
