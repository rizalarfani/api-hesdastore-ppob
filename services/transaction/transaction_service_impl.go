package services

import (
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"strings"

	"hesdastore/api-ppob/clients/config"
	clients "hesdastore/api-ppob/clients/digiflazz"
	"hesdastore/api-ppob/common/helper"
	"hesdastore/api-ppob/constants"
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

func (s *TransactionServiceImpl) GetHistory(ctx context.Context, trxID string, userId int) ([]*dto.TransactionHistoryResponse, error) {
	historys, err := s.repository.Transaction().GetAll(ctx, trxID, userId)
	if err != nil {
		return nil, err
	}

	data := make([]*dto.TransactionHistoryResponse, 0, len(historys))
	for _, p := range historys {
		data = append(data, &dto.TransactionHistoryResponse{
			TransactionsID: p.TransactionID,
			ProductName:    p.PackageName,
			Brand: &dto.BrandResponse{
				Name: p.Brand.Name,
				Logo: constants.UploadBrandUrl + "/" + p.Brand.Logo,
			},
			Category: &dto.CategoryResponse{
				Name: p.Category.Name,
			},
			Price:      p.Price,
			CustomerNo: p.PhoneNumber,
			SN:         p.SN,
			Status:     p.Status.GetStatusString(),
			Message:    p.StatusMessage,
		})
	}
	return data, nil
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

	secret, err := s.repository.AuthApi().FindSecretKeyByUserID(ctx, auth.UserID)
	if err != nil {
		return nil, err
	}

	signature := helper.GenerateTransactionSignature(
		secret,
		refID,
		request.CustomerNo,
		request.ProductCode,
		price,
	)

	order, err = s.repository.Transaction().CreateOrder(ctx, tx, &model.TransactionOrder{
		UserID:          auth.UserID,
		PackageID:       product.ProductID,
		PackageCode:     product.ProductID,
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
		Signature:       &signature,
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
		SN:             topupResponse.Data.Sn,
	}
	return &response, nil
}

func (s *TransactionServiceImpl) Webhooks(
	ctx context.Context,
	headerSignature string,
	payload []byte,
) error {
	if !verifySignature(headerSignature, payload) {
		return errors.New("invalid validate signature")
	}

	var req dto.DigifalzzWebhooksPayload
	if err := json.Unmarshal(payload, &req); err != nil {
		return err
	}

	order, err := s.repository.Transaction().GetTransactionByRefID(ctx, req.Data.RefID)
	if err != nil {
		return err
	}

	tx, err := s.repository.GetTx().BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	calbackResponse, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	switch req.Data.Status {
	case "Sukses":
		err = s.repository.Transaction().UpdateTransaction(ctx, tx, &dto.TransactionUpdateRequest{
			TrxId:     order.TransactionID,
			Response:  string(calbackResponse),
			Status:    1,
			StatusMsg: req.Data.Message,
			Sn:        *req.Data.Sn,
		})
		if err != nil {
			return nil
		}
	case "Gagal":
		err = s.repository.Transaction().UpdateTransaction(ctx, tx, &dto.TransactionUpdateRequest{
			TrxId:     order.TransactionID,
			Response:  string(calbackResponse),
			Status:    2,
			StatusMsg: req.Data.Message,
			Sn:        *req.Data.Sn,
		})

		if err != nil {
			return nil
		}

		balance, err := s.repository.Account().FindBalanceUserByUserId(ctx, order.UserID)
		if err != nil {
			return nil
		}

		newBalance := balance.Balance + order.Price
		err = s.repository.Transaction().UpdateBalance(ctx, tx, &dto.TransactionUpdateBalanceRequest{
			UserID:     order.UserID,
			NewBalance: newBalance,
		})
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
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

func verifySignature(signatureHeader string, rawBody []byte) bool {
	mac := hmac.New(sha1.New, []byte(helper.GetEnv("SECRET_KEY_DIGIFLAZZ")))
	mac.Write(rawBody)
	expectSignature := hex.EncodeToString(mac.Sum(nil))

	parts := strings.SplitN(signatureHeader, "=", 2)
	if len(parts) != 2 || parts[0] != "sha1" {
		return false
	}
	return hmac.Equal([]byte(parts[1]), []byte(expectSignature))
}
