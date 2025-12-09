package services

import (
	"context"
	"encoding/json"
	"fmt"
	"hesdastore/api-ppob/clients/config"
	clients "hesdastore/api-ppob/clients/digiflazz"
	"hesdastore/api-ppob/common/helper"
	"hesdastore/api-ppob/constants"
	"hesdastore/api-ppob/domain/dto"
	"hesdastore/api-ppob/domain/model"
	"hesdastore/api-ppob/repositories"

	errConstant "hesdastore/api-ppob/constants/error"
)

type BillingServiceImpl struct {
	repository repositories.IRepoRegistry
	digifalzz  clients.IDigiflazzClient
	client     config.IClientConfig
}

func NewBillingServiceImpl(repo repositories.IRepoRegistry, digifalzz clients.IDigiflazzClient,
	client config.IClientConfig) BillingService {
	return &BillingServiceImpl{
		repository: repo,
		digifalzz:  digifalzz,
		client:     client,
	}
}

func (s *BillingServiceImpl) Inquiry(
	ctx context.Context,
	request *dto.InquiryBillRequest,
	auth *model.ApiUser,
) (*dto.InquiryBillingResponse, error) {
	var (
		err     error
		balance *model.Account
		product *model.Product
	)

	balance, err = s.repository.Account().FindBalanceUser(ctx, auth.Username)
	if err != nil {
		return nil, err
	}

	if balance == nil {
		return nil, errConstant.ErrInternalServerError
	}
	if balance.Balance == 0 {
		return nil, errConstant.ErrBalanceIsZero
	}

	product, err = s.repository.Product().FindByProductCode(ctx, request.ProductCode)
	if err != nil {
		return nil, err
	}

	status := helper.GetStatusProductByRole(auth.Role, *product)
	if status == "inactive" {
		return nil, errConstant.ErrProductIsAvalaible
	}

	refID, err := helper.GenerateRandomString(15)
	if err != nil {
		return nil, err
	}

	inquiry := &clients.InquiryRequest{
		CustomerNo: request.CustomerNo,
		SKUCode:    request.ProductCode,
		RefID:      refID,
	}

	inquiryResponse, err := s.digifalzz.Inquiry(ctx, inquiry)
	if err != nil {
		return nil, err
	}

	tx, err := s.repository.GetTx().BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	inquiryResString, err := json.Marshal(inquiryResponse)
	if err != nil {
		return nil, err
	}

	create, errDb := s.repository.Billing().CreateInquiry(ctx, tx, &model.InquiryBilling{
		UserID:          auth.UserID,
		PackageID:       product.ProductID,
		PackageCode:     product.ProductID,
		PackageName:     product.ProductName,
		TransactionID:   inquiryResponse.Data.RefID,
		PhoneNumber:     request.CustomerNo,
		OriginalPrice:   inquiryResponse.Data.OriginalPrice,
		Price:           inquiryResponse.Data.Price,
		FinalBalance:    0,
		NewBalance:      0,
		Status:          4, // Inquiry
		StatusMessage:   inquiryResponse.Data.Message,
		Type:            "digiflazz",
		TransactionFrom: "api",
		Response:        string(inquiryResString),
	})
	if errDb != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	response := dto.InquiryBillingResponse{
		TransactionsID: create.TransactionID,
		ProductCode:    request.ProductCode,
		Brand: &dto.BrandResponse{
			Name: product.Brand.Name,
			Logo: constants.UploadBrandUrl + "/" + product.Brand.Logo,
		},
		Category: &dto.CategoryResponse{
			Name: product.Category.Name,
		},
		CustomerName: inquiryResponse.Data.CustomerName,
		Price:        inquiryResponse.Data.Price,
		Status:       create.Status.GetStatusString(),
		Message:      create.StatusMessage,
	}

	switch product.Category.ID {
	case 11:
		var internetDesc dto.InternetDescRespDigiflazz
		if err := json.Unmarshal([]byte(inquiryResponse.Data.Desc), &internetDesc); err != nil {
			return nil, fmt.Errorf("failed unmarshal digiflazz response: %w", err)
		}

		data := make([]dto.InternetDetailResponse, 0, len(internetDesc.Detail))
		for _, b := range internetDesc.Detail {
			data = append(data, dto.InternetDetailResponse{
				Period:     b.Period,
				BillValue:  b.BillValue,
				PriceAdmin: b.PriceAdmin,
			})
		}

		response.Detail = &dto.InternetDescResponse{
			BillingSheet: internetDesc.BillingSheet,
			Detail:       data,
		}
	case 9:
		var plnDesc dto.PlnDescRespDigiflazz
		if err := json.Unmarshal([]byte(inquiryResponse.Data.Desc), &plnDesc); err != nil {
			return nil, fmt.Errorf("failed unmarshal digiflazz response: %w", err)
		}

		data := make([]dto.PlnDetailResponse, 0, len(plnDesc.Detail))
		for _, b := range plnDesc.Detail {
			data = append(data, dto.PlnDetailResponse{
				Period:     b.Period,
				BillValue:  b.BillValue,
				PriceAdmin: b.PriceAdmin,
				Fine:       b.Fine,
			})
		}

		response.Detail = &dto.PlnDescResponse{
			Tarif:        plnDesc.Tarif,
			Power:        plnDesc.Power,
			BillingSheet: plnDesc.BillingSheet,
			Detail:       data,
		}
	case 10:
		var pdamDesc dto.PdamDescRespDigiflazz
		if err := json.Unmarshal([]byte(inquiryResponse.Data.Desc), &pdamDesc); err != nil {
			return nil, fmt.Errorf("failed unmarshal digiflazz response: %w", err)
		}

		data := make([]dto.PdamDetailResponse, 0, len(pdamDesc.Detail))
		for _, b := range pdamDesc.Detail {
			data = append(data, dto.PdamDetailResponse{
				Period:    b.Period,
				BillValue: b.BillValue,
				Fine:      b.Fine,
			})
		}

		response.Detail = &dto.PdamDescResponse{
			Tarif:        pdamDesc.Tarif,
			BillingSheet: pdamDesc.BillingSheet,
			Address:      pdamDesc.Address,
			DueDate:      pdamDesc.DueDate,
			Detail:       data,
		}
	case 12:
		var bpjsDesc dto.BpjsKesDescRespDigiflazz
		if err := json.Unmarshal([]byte(inquiryResponse.Data.Desc), &bpjsDesc); err != nil {
			return nil, fmt.Errorf("failed unmarshal digiflazz response: %w", err)
		}

		data := make([]dto.BpjsKesDetailResponse, 0, len(bpjsDesc.Detail))
		for _, b := range bpjsDesc.Detail {
			data = append(data, dto.BpjsKesDetailResponse{
				Period: b.Period,
			})
		}

		response.Detail = &dto.BpjsKesDescResponse{
			BillingSheet: bpjsDesc.BillingSheet,
			Address:      bpjsDesc.Address,
			Detail:       data,
		}
	}

	return &response, nil
}
