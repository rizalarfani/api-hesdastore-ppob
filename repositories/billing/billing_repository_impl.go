package repositories

import (
	"context"
	errWrap "hesdastore/api-ppob/common/error"
	errConstant "hesdastore/api-ppob/constants/error"
	"hesdastore/api-ppob/domain/model"
	"log"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type BillingRepositoryImpl struct {
	db *sqlx.DB
	qb squirrel.StatementBuilderType
}

func NewBillingRepositoryImpl(db *sqlx.DB) BillingRepository {
	return &BillingRepositoryImpl{
		db: db,
		qb: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Question),
	}
}

func (r *BillingRepositoryImpl) CreateInquiry(ctx context.Context, tx *sqlx.Tx, order *model.InquiryBilling) (*model.InquiryBilling, error) {
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
		)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, errWrap.WrapError(errConstant.ErrSQLError)
	}

	res, err := tx.ExecContext(ctx, query, args...)
	log.Println(err)
	if err != nil {
		return nil, err
	}

	if lastID, err := res.LastInsertId(); err == nil {
		_ = lastID
	}

	return order, nil
}
