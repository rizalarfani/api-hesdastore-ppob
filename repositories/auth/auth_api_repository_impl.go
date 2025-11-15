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

type AuthApiRepositoryImpl struct {
	db *sqlx.DB
	qb squirrel.StatementBuilderType
}

func NewAuthApiRepositoryImpl(db *sqlx.DB) AuhtApiRepository {
	return &AuthApiRepositoryImpl{
		db: db,
		qb: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Question),
	}
}

func (r *AuthApiRepositoryImpl) FindByUsername(ctx context.Context, username string) (*model.ApiUser, error) {
	builder := r.qb.Select("username,password,`keys`.`key`").
		From("users").
		Join("`keys` ON users.id = `keys`.user_id").
		Where(squirrel.Eq{
			"users.username": username,
		}).
		Limit(1)
	sqlStr, args, err := builder.ToSql()

	if err != nil {
		return nil, errWrap.WrapError(err)
	}

	var user model.ApiUser
	if err := r.db.GetContext(ctx, &user, sqlStr, args...); err != nil {
		if err == sql.ErrNoRows {
			return nil, errConstant.ErrUserNotFound
		}
		return nil, errConstant.ErrInternalServerError
	}

	return &user, nil
}
