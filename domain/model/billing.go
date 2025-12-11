package model

import "hesdastore/api-ppob/constants"

type InquiryBilling struct {
	UserID          int                         `db:"idUser"`
	PackageID       int                         `db:"id_paket"`
	PackageCode     int                         `db:"package_code"`
	PackageName     string                      `db:"nama_paket"`
	TransactionID   string                      `db:"trx_id"`
	PhoneNumber     string                      `db:"no_hp"`
	Response        string                      `db:"res"`
	OriginalPrice   int                         `db:"harga_asli"`
	Price           int                         `db:"harga"`
	FinalBalance    int                         `db:"saldo_akhir"`
	NewBalance      int                         `db:"saldo_baru"`
	Status          constants.TransactionStatus `db:"status"`
	StatusMessage   string                      `db:"status_msg"`
	Type            string                      `db:"tipe"`
	TransactionFrom string                      `db:"trx_from"`
}

type PayBilling struct {
	TransactionID string                      `db:"trx_id"`
	Response      string                      `db:"res"`
	FinalBalance  int                         `db:"saldo_akhir"`
	NewBalance    int                         `db:"saldo_baru"`
	Status        constants.TransactionStatus `db:"status"`
	StatusMessage string                      `db:"status_msg"`
	SN            string                      `db:"sn"`
	CallbackURL   *string                     `db:"url_callback"`
	Signature     *string                     `db:"signature"`
}
