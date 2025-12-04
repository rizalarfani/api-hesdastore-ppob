package clients

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"hesdastore/api-ppob/clients/config"
	"hesdastore/api-ppob/constants"

	errConstant "hesdastore/api-ppob/constants/error"
)

type DigiflazzClient struct {
	client config.IClientConfig
}

type IDigiflazzClient interface {
	Topup(context.Context, *TopupRequest) (*TopupResponse, error)
}

func NewDigiflazzClient(client config.IClientConfig) *DigiflazzClient {
	return &DigiflazzClient{
		client: client,
	}
}

func (c *DigiflazzClient) Topup(ctx context.Context, req *TopupRequest) (*TopupResponse, error) {
	dataSign := []byte(c.client.Username() + c.client.ApiKey() + req.RefID)
	signature := md5.Sum(dataSign)
	req.Username = c.client.Username()
	req.Signature = hex.EncodeToString(signature[:])
	req.CalbackURL = constants.DigiflazzWebhooksUrl

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	_, bodyResp, errs := c.client.Client().
		Post(fmt.Sprintf("%s/transaction", c.client.BaseURL())).
		Send(string(body)).
		End()

	if len(errs) > 0 {
		return nil, errs[0]
	}

	var response TopupResponse
	if err := json.Unmarshal([]byte(bodyResp), &response); err != nil {
		return nil, fmt.Errorf("failed unmarshal digiflazz response: %w", err)
	}

	switch response.Data.Rc {
	case "44", "42", "69":
		return nil, errConstant.ErrServiceNotAvailable
	}

	if response.Data.Rc != "03" {
		return nil, errors.New(response.Data.Message)
	}

	return &response, nil
}
