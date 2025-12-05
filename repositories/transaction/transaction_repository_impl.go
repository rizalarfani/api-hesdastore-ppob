package repositories

import (
	"context"
	"database/sql"
	errWrap "hesdastore/api-ppob/common/error"
	errConstant "hesdastore/api-ppob/constants/error"
	"hesdastore/api-ppob/domain/dto"
	"hesdastore/api-ppob/domain/model"
	"log"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type TransactionRepositoryImpl struct {
	db *sqlx.DB
	qb squirrel.StatementBuilderType
}

func NewProductRepositoryImpl(db *sqlx.DB) TransactionRepository {
	return &TransactionRepositoryImpl{
		db: db,
		qb: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Question),
	}
}

func (r *TransactionRepositoryImpl) CreateOrder(ctx context.Context, tx *sqlx.Tx, order *model.TransactionOrder) (*model.TransactionOrder, error) {
	builder := r.qb.Insert("transaksi").
		Columns(
			"idUser",
			"id_paket",
			"package_code",
			"nama_paket",
			"trx_id",
			"no_hp",
			"res",
			"harga_asli",
			"harga",
			"saldo_akhir",
			"saldo_baru",
			"status",
			"status_msg",
			"tipe",
			"trx_from",
			"url_callback",
			"signature",
		).
		Values(
			order.UserID,
			order.PackageID,
			order.PackageCode,
			order.PackageName,
			order.TransactionID,
			order.PhoneNumber,
			order.Response,
			order.OriginalPrice,
			order.Price,
			order.FinalBalance,
			order.NewBalance,
			order.Status,
			order.StatusMessage,
			order.Type,
			order.TransactionFrom,
			order.CallbackURL,
			order.Signature,
		)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, errWrap.WrapError(errConstant.ErrSQLError)
	}

	res, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	if lastID, err := res.LastInsertId(); err == nil {
		_ = lastID
	}

	return order, nil
}

func (r *TransactionRepositoryImpl) UpdateBalance(ctx context.Context, tx *sqlx.Tx, req *dto.TransactionUpdateBalanceRequest) error {
	builder := r.qb.Update("users").
		Set("saldo", req.NewBalance).
		Where(squirrel.Eq{
			"id": req.UserID,
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return errWrap.WrapError(errConstant.ErrSQLError)
	}

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *TransactionRepositoryImpl) GetTransactionByRefID(ctx context.Context, refID string) (*model.TransactionOrder, error) {
	builder := r.qb.Select(`
		idUser,
		id_paket,
		package_code,
		nama_paket,
		trx_id,
		no_hp,
		res,
		harga_asli,
		harga,
		saldo_akhir,
		saldo_baru,
		status,
		status_msg,
		tipe,
		trx_from,
		url_callback,
		signature
	`).From("transaksi").Where(squirrel.Eq{
		"transaksi.trx_id":   refID,
		"transaksi.trx_from": "api",
	})

	strSql, args, err := builder.ToSql()
	if err != nil {
		return nil, errWrap.WrapError(err)
	}

	var order model.TransactionOrder
	if err := r.db.GetContext(ctx, &order, strSql, args...); err != nil {
		if err == sql.ErrNoRows {
			return nil, errConstant.ErrTransactionNotFound
		}
		return nil, errConstant.ErrInternalServerError
	}

	return &order, nil
}

func (r *TransactionRepositoryImpl) UpdateTransaction(ctx context.Context, tx *sqlx.Tx, req *dto.TransactionUpdateRequest) error {
	builder := r.qb.Update("transaksi").
		Set("res", req.Response).
		Set("status", req.Status).
		Set("status_msg", req.StatusMsg).
		Set("sn", req.Sn).
		Where(squirrel.Eq{
			"trx_id": req.TrxId,
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return errWrap.WrapError(errConstant.ErrSQLError)
	}

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *TransactionRepositoryImpl) GetAll(ctx context.Context, trxID string, userID int) ([]*model.TransactionHistory, error) {
	builder := r.qb.Select(`
		nama_paket,
		trx_id,
		no_hp,
		harga,
		transaksi.status,
		status_msg,
		category.name as category_name,
		category.id as category_id,
		brands.name,
		brands.logo,
		transaksi.sn
	`).
		From("transaksi").
		Join("products ON products.id = transaksi.id_paket").
		Join("category ON category.id = products.category_id").
		Join("brands ON brands.id = products.brand_id").
		Where(squirrel.Eq{
			"transaksi.idUser":   userID,
			"transaksi.trx_from": "api",
		}).
		OrderBy("transaksi.id desc")

	if trxID != "" {
		builder = builder.Where(squirrel.Eq{
			"transaksi.trx_id": trxID,
		})
	}

	strSql, args, err := builder.ToSql()
	if err != nil {
		return nil, errWrap.WrapError(err)
	}

	var rows []*model.TransactionHistory
	if err := r.db.SelectContext(ctx, &rows, strSql, args...); err != nil {
		log.Println(err)
		if err == sql.ErrNoRows {
			return nil, errConstant.ErrTransactionNotFound
		}

		return nil, errConstant.ErrInternalServerError
	}

	return rows, nil
}
