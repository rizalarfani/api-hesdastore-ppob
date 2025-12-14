package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"hesdastore/api-ppob/domain/dto"
	"hesdastore/api-ppob/domain/model"
	"hesdastore/api-ppob/repositories"
	"io"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type WebhookServiceImpl struct {
	repository repositories.IRepoRegistry
	httpClient *http.Client
	log        *logrus.Logger
}

func NewWebhookServiceImpl(
	repository repositories.IRepoRegistry,
) WebhookService {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	return &WebhookServiceImpl{
		repository: repository,
		httpClient: &http.Client{
			Timeout: time.Second * 30,
		},
		log: logger,
	}
}

func (s *WebhookServiceImpl) SendWebhook(ctx context.Context, payload *dto.TransactionUpdateEventWebhook) error {
	delevery_ref := payload.TransactionID

	s.log.WithFields(logrus.Fields{
		"delivery_ref": delevery_ref,
		"callback_url": payload.CallbackURL,
	}).Info("Send webhook to client")

	clientPayload := dto.WebhooksPayloadToClient{
		TransactionID: payload.TransactionID,
		ProductName:   payload.ProductName,
		CustomerNo:    payload.CustomerNo,
		Status:        payload.Status,
		StatusMessage: payload.StatusMessage,
		SN:            payload.SN,
		Timestamp:     time.Now().Format(time.RFC3339),
	}

	jsonPayload, err := json.Marshal(clientPayload)
	if err != nil {
		return s.logError(delevery_ref, payload, "", http.StatusInternalServerError, err)
	}

	// create HTTP reqeuest to client
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		payload.CallbackURL,
		bytes.NewBuffer(jsonPayload),
	)

	if err != nil {
		return s.logError(delevery_ref, payload, string(jsonPayload), 0, err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hesda-Event", payload.EventType)
	req.Header.Set("X-Hesda-Signature", payload.Signature)

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return s.logError(delevery_ref, payload, string(jsonPayload), 0, err)
	}
	defer resp.Body.Close()

	responseBody, _ := io.ReadAll(resp.Body)
	var responError string

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		responError = fmt.Sprintf("HTTP %d: %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	} else {
		responError = ""
	}

	err = s.repository.Webhook().Create(context.Background(), &model.Webhook{
		EventType:      payload.EventType,
		DeleveryRef:    delevery_ref,
		Endpoint:       payload.CallbackURL,
		RequestBody:    string(jsonPayload),
		ResponseBody:   string(responseBody),
		ResponseStatus: resp.StatusCode,
		ResponseError:  responError,
		Signature:      payload.Signature,
	})
	return err
}

func (s *WebhookServiceImpl) logError(
	deliveryRef string,
	payload *dto.TransactionUpdateEventWebhook,
	requestBody string,
	statusCode int,
	err error,
) error {
	s.log.WithFields(logrus.Fields{
		"delivery_ref":   deliveryRef,
		"transaction_id": payload.TransactionID,
		"error":          err.Error(),
	}).Error("Webhook failed")

	_ = s.repository.Webhook().Create(context.Background(), &model.Webhook{
		EventType:      "transaction.update",
		DeleveryRef:    deliveryRef,
		Endpoint:       payload.CallbackURL,
		RequestBody:    requestBody,
		ResponseBody:   "",
		ResponseStatus: statusCode,
		ResponseError:  err.Error(),
		Signature:      payload.Signature,
	})

	return err
}
