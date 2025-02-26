package transaction

import "golang-voucher-api/internal/voucher"

type Transaction struct {
	ID         uint            `gorm:"primaryKey" json:"id"`
	VoucherID  uint            `gorm:"not null" json:"voucher_id"`
	Quantity   int             `gorm:"not null" json:"quantity"`
	TotalPoint int             `gorm:"not null" json:"total_point"`
	Voucher    voucher.Voucher `gorm:"foreignKey:VoucherID" json:"voucher"`
}
