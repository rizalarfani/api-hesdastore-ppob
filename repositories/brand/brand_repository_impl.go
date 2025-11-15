package repositories

import (
	"context"
	errWrap "hesdastore/api-ppob/common/error"
	"hesdastore/api-ppob/domain/model"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type BrandRepositoryImpl struct {
	db *sqlx.DB
	qb squirrel.StatementBuilderType
}

func NewBrandRepositoryImpl(db *sqlx.DB) BrandRepository {
	return &BrandRepositoryImpl{
		db: db,
		qb: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Question),
	}
}

func (r *BrandRepositoryImpl) FindAll(ctx context.Context) ([]*model.Brand, error) {
	builder := r.qb.Select("id,name,logo").From("brands").OrderBy("name ASC")

	sqlStr, args, err := builder.ToSql()
	if err != nil {
		return nil, errWrap.WrapError(err)
	}

	rows, err := r.db.QueryContext(ctx, sqlStr, args...)
	if err != nil {
		return nil, errWrap.WrapError(err)
	}
	defer rows.Close()

	var brands []*model.Brand
	for rows.Next() {
		brand := &model.Brand{}
		err := rows.Scan(&brand.ID, &brand.Name, &brand.Logo)
		if err != nil {
			panic(err)
		}

		brands = append(brands, brand)
	}

	return brands, nil
}
