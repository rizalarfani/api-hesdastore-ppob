package repositories

import (
	"context"
	errWrap "hesdastore/api-ppob/common/error"
	errConstant "hesdastore/api-ppob/constants/error"
	"hesdastore/api-ppob/domain/model"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type WebhookRepositoryImpl struct {
	db *sqlx.DB
	qb squirrel.StatementBuilderType
}

func NewWebhookRepositoryImpl(db *sqlx.DB) WebhookRepository {
	return &WebhookRepositoryImpl{
		db: db,
		qb: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Question),
	}
}

func (r *WebhookRepositoryImpl) Create(ctx context.Context, webhook *model.Webhook) error {
	builder := r.qb.Insert("webhook_logs").
		Columns("event_type", "delivery_ref", "endpoint", "request_body", "response_body", "response_status", "response_error", "signature").
		Values(webhook.EventType, webhook.DeleveryRef, webhook.Endpoint, webhook.RequestBody, webhook.ResponseBody, webhook.ResponseStatus, webhook.ResponseError, webhook.Signature)

	query, args, err := builder.ToSql()
	if err != nil {
		return errWrap.WrapError(errConstant.ErrSQLError)
	}

	_, err = r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *WebhookRepositoryImpl) GetDeliveryRef(ctx context.Context, deleveryRef string) (*model.Webhook, error) {
	builder := r.qb.Select("*").
		From("webhook_logs").
		Where(squirrel.Eq{"delivery_ref": deleveryRef})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, errWrap.WrapError(errConstant.ErrSQLError)
	}

	var webhook model.Webhook
	err = r.db.GetContext(ctx, &webhook, query, args...)
	if err != nil {
		return nil, errWrap.WrapError(errConstant.ErrSQLError)
	}

	return &webhook, nil
}
