package services

import (
	"context"
	"encoding/json"

	"hesdastore/api-ppob/clients/config"
	clients "hesdastore/api-ppob/clients/digiflazz"
	"hesdastore/api-ppob/common/helper"
	"hesdastore/api-ppob/domain/dto"
	"hesdastore/api-ppob/domain/model"
	"hesdastore/api-ppob/repositories"

	errConstant "hesdastore/api-ppob/constants/error"
)

type TransactionServiceImpl struct {
	repository repositories.IRepoRegistry
	digifalzz  clients.IDigiflazzClient
	client     config.IClientConfig
}

func NewTransactionServiceImpl(
	repo repositories.IRepoRegistry,
	digifalzz clients.IDigiflazzClient,
	client config.IClientConfig,
) TransactionService {
	return &TransactionServiceImpl{
		repository: repo,
		digifalzz:  digifalzz,
	}
}

func (s *TransactionServiceImpl) Order(
	ctx context.Context,
	request *dto.TransactionOrderRequest,
	auth *model.ApiUser,
) (*dto.TransactionOrderResponse, error) {
	var (
		topupErr, err error
		balance       *model.Account
		product       *model.Product
		price         int
		topupReq      *clients.TopupRequest
		topupResponse *clients.TopupResponse
		order         *model.TransactionOrder
	)

	balance, err = s.repository.Account().FindBalanceUser(ctx, auth.Username)
	if err != nil {
		return nil, err
	}

	if err := s.validateBalance(balance); err != nil {
		return nil, err
	}

	product, err = s.repository.Product().FindByProductCode(ctx, request.ProductCode)
	if err != nil {
		return nil, err
	}

	price, err = s.getProductPriceAndValidate(auth.Role, product)
	if err != nil {
		return nil, err
	}

	if balance.Balance < price {
		return nil, errConstant.ErrBalanceIsNotEnough
	}

	if helper.InIntSlice(product.Category.ID, []int{1, 3, 6}) {
		request.CustomerNo = helper.ToLocal08(request.CustomerNo)
	}

	refID, err := helper.GenerateRandomString(15)
	if err != nil {
		return nil, err
	}

	topupReq = &clients.TopupRequest{
		SKUCode:    product.ProductCode,
		CustomerNo: request.CustomerNo,
		RefID:      refID,
	}

	topupResponse, topupErr = s.digifalzz.Topup(ctx, topupReq)
	if topupErr != nil {
		return nil, topupErr
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

	newBalance := balance.Balance - price
	topupResString, err := json.Marshal(topupResponse)
	if err != nil {
		return nil, err
	}

	order, err = s.repository.Transaction().CreateOrder(ctx, tx, &model.TransactionOrder{
		UserID:          auth.UserID,
		PackageID:       product.ProductID,
		PackageCode:     request.ProductCode,
		PackageName:     product.ProductName,
		TransactionID:   topupResponse.Data.RefID,
		PhoneNumber:     request.CustomerNo,
		OriginalPrice:   *product.Hpp,
		Price:           price,
		FinalBalance:    balance.Balance,
		NewBalance:      newBalance,
		Status:          0, // PENDING
		StatusMessage:   topupResponse.Data.Message,
		Type:            "digiflazz",
		TransactionFrom: "api",
		CallbackURL:     &request.CallbackURL,
		Response:        string(topupResString),
	})
	if err != nil {
		return nil, err
	}

	err = s.repository.Transaction().UpdateBalance(ctx, tx, &dto.TransactionUpdateBalanceRequest{
		UserID:     auth.UserID,
		NewBalance: newBalance,
	})
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	response := dto.TransactionOrderResponse{
		TransactionsID: order.TransactionID,
		ProductCode:    product.ProductCode,
		ProductName:    product.ProductName,
		Status:         order.Status.GetStatusString(),
		Message:        order.StatusMessage,
	}
	return &response, nil
}

func (s *TransactionServiceImpl) validateBalance(balance *model.Account) error {
	if balance == nil {
		return errConstant.ErrInternalServerError
	}
	if balance.Balance == 0 {
		return errConstant.ErrBalanceIsZero
	}
	return nil
}

func (s *TransactionServiceImpl) getProductPriceAndValidate(role int, product *model.Product) (int, error) {
	if product == nil {
		return 0, errConstant.ErrProductIsFaulty
	}

	if !s.isProductActive(role, product) {
		return 0, errConstant.ErrProductIsAvalaible
	}

	if product.Hpp == nil {
		return 0, errConstant.ErrProductIsFaulty
	}

	price := helper.GetPriceProductByRole(role, *product)
	if *product.Hpp > price {
		return 0, errConstant.ErrProductIsFaulty
	}

	return price, nil
}

func (s *TransactionServiceImpl) isProductActive(role int, product *model.Product) bool {
	status := helper.GetStatusProductByRole(role, *product)
	return status != "inactive"
}
