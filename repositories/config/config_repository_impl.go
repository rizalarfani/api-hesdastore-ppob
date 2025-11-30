package repositories

import (
	"context"
	errWrap "hesdastore/api-ppob/common/error"
	errConstant "hesdastore/api-ppob/constants/error"
	"hesdastore/api-ppob/domain/model"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type ConfigRepositoryImpl struct {
	db *sqlx.DB
	qb squirrel.StatementBuilderType
}

func NewConfigRepositoryImpl(db *sqlx.DB) ConfigRepository {
	return &ConfigRepositoryImpl{
		db: db,
		qb: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Question),
	}
}

func (r *ConfigRepositoryImpl) GetConfigDigiflazz(ctx context.Context) (*model.Digiflazz, error) {
	builder := r.qb.Select("url_digiflaz,api_key_digiflaz,username_digiflaz").
		From("config").
		Where(squirrel.Eq{
			"id": 1,
		}).
		Limit(1)

	strSql, args, err := builder.ToSql()
	if err != nil {
		return nil, errWrap.WrapError(err)
	}

	var configDigiflazz model.Digiflazz
	if err := r.db.GetContext(ctx, &configDigiflazz, strSql, args...); err != nil {
		return nil, errConstant.ErrInternalServerError
	}

	return &configDigiflazz, nil
}
