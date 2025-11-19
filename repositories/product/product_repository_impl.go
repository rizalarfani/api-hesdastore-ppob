package repositories

import (
	"context"
	"database/sql"
	errWrap "hesdastore/api-ppob/common/error"
	errConstant "hesdastore/api-ppob/constants/error"
	"hesdastore/api-ppob/domain/model"
	"log"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type ProductRepositoryImpl struct {
	db *sqlx.DB
	qb squirrel.StatementBuilderType
}

func NewProductRepositoryImpl(db *sqlx.DB) ProductRepository {
	return &ProductRepositoryImpl{
		db: db,
		qb: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Question),
	}
}

func (r *ProductRepositoryImpl) FindAllPrabayar(ctx context.Context) ([]*model.Product, error) {
	sqlBuilder := r.qb.Select(`
		products.metode,
		products.product_code,
		products.product_name,
		category.name as category_name,
		brands.name,
		brands.logo,
		products.type,
		products.price_seller,
		products.price_reseller,
		products.seller_status,
		products.reseller_status,
		products.start_cut_off,
		products.end_cut_off,
		products.description
	`).
		From("products").
		Join("category ON category.id = products.category_id").Join("brands ON brands.id = products.brand_id").
		Where(squirrel.Eq{
			"products.is_deleted": 0,
			"products.metode":     "prepaid",
		}).
		OrderBy("products.price_seller asc").
		OrderBy("products.price_reseller asc")

	strSql, args, err := sqlBuilder.ToSql()
	if err != nil {
		return nil, errWrap.WrapError(err)
	}

	var rows []*model.Product
	if err := r.db.SelectContext(ctx, &rows, strSql, args...); err != nil {
		if err == sql.ErrNoRows {
			return nil, errConstant.ErrProductNotFound
		}

		return nil, errConstant.ErrInternalServerError
	}

	return rows, nil
}

func (r *ProductRepositoryImpl) FindByProductCode(ctx context.Context, productCode string) (*model.Product, error) {
	sqlBuilder := r.qb.Select(`
		products.metode,
		products.product_code,
		products.product_name,
		category.name as category_name,
		brands.name,
		brands.logo,
		products.type,
		products.price_seller,
		products.price_reseller,
		products.admin,
		products.komisi as commission,
		products.seller_status,
		products.reseller_status,
		products.start_cut_off,
		products.end_cut_off,
		products.description
	`).
		From("products").
		Join("category ON category.id = products.category_id").Join("brands ON brands.id = products.brand_id").
		Where(squirrel.Eq{
			"products.is_deleted":   0,
			"products.product_code": productCode,
		}).
		OrderBy("products.price_seller asc").
		OrderBy("products.price_reseller asc")

	strSql, args, err := sqlBuilder.ToSql()
	if err != nil {
		return nil, errWrap.WrapError(err)
	}

	var product model.Product
	if err := r.db.GetContext(ctx, &product, strSql, args...); err != nil {
		if err == sql.ErrNoRows {
			return nil, errConstant.ErrProductNotFound
		}
		log.Println(err)
		return nil, errConstant.ErrInternalServerError
	}

	return &product, nil
}

func (r *ProductRepositoryImpl) FindAllPascabayar(ctx context.Context) ([]*model.Product, error) {
	sqlBuilder := r.qb.Select(`
		products.metode,
		products.product_code,
		products.product_name,
		category.name as category_name,
		brands.name,
		brands.logo,
		products.type,
		products.admin,
		products.komisi as commission,
		products.seller_status,
		products.reseller_status,		
		products.description
	`).
		From("products").
		Join("category ON category.id = products.category_id").
		Join("brands ON brands.id = products.brand_id").
		Where(squirrel.Eq{
			"products.is_deleted": 0,
			"products.metode":     "pasca",
		}).
		OrderBy("products.price_seller asc").
		OrderBy("products.price_reseller asc")

	strSql, args, err := sqlBuilder.ToSql()
	if err != nil {
		return nil, errWrap.WrapError(err)
	}

	var rows []*model.Product
	if err := r.db.SelectContext(ctx, &rows, strSql, args...); err != nil {
		if err == sql.ErrNoRows {
			return nil, errConstant.ErrProductNotFound
		}

		return nil, errConstant.ErrInternalServerError
	}

	return rows, nil
}
