package repositories

import (
	"context"
	"database/sql"
	errWrap "hesdastore/api-ppob/common/error"
	errConstant "hesdastore/api-ppob/constants/error"
	"hesdastore/api-ppob/domain/model"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type AccountRepositoryImpl struct {
	db *sqlx.DB
	qb squirrel.StatementBuilderType
}

func NewAccountRepositoryImpl(db *sqlx.DB) AccountRepository {
	return &AccountRepositoryImpl{
		db: db,
		qb: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Question),
	}
}

func (r *AccountRepositoryImpl) FindBalanceUser(ctx context.Context, username string) (*model.Account, error) {
	sqlBuilder := r.qb.Select("nama,saldo").
		From("users").
		Where(squirrel.Eq{
			"users.username": username,
		}).
		Limit(1)

	strSql, args, err := sqlBuilder.ToSql()
	if err != nil {
		return nil, errWrap.WrapError(err)
	}

	var account model.Account
	if err := r.db.GetContext(ctx, &account, strSql, args...); err != nil {
		if err == sql.ErrNoRows {
			return nil, errConstant.ErrUserNotFound
		}

		return nil, errConstant.ErrInternalServerError
	}

	return &account, nil
}

func (r *AccountRepositoryImpl) FindBalanceUserByUserId(ctx context.Context, userId int) (*model.Account, error) {
	sqlBuilder := r.qb.Select("nama,saldo").
		From("users").
		Where(squirrel.Eq{
			"users.id": userId,
		}).
		Limit(1)

	strSql, args, err := sqlBuilder.ToSql()
	if err != nil {
		return nil, errWrap.WrapError(err)
	}

	var account model.Account
	if err := r.db.GetContext(ctx, &account, strSql, args...); err != nil {
		if err == sql.ErrNoRows {
			return nil, errConstant.ErrUserNotFound
		}

		return nil, errConstant.ErrInternalServerError
	}

	return &account, nil
}
