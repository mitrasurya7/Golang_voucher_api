package transaction

import (
	"errors"
	"fmt"
	"golang-voucher-api/internal/voucher"
)

func CreateNewTransaction(voucherID uint, quantity int, totalPoint int) (*Transaction, error) {
	v, err := voucher.GetSingleVoucherByID(voucherID)
	if err != nil {
		return nil, fmt.Errorf("gagal mendapatkan voucher dengan ID %d: %w", voucherID, err)
	}

	if v.CostInPoint*quantity != totalPoint {
		return nil, errors.New("total poin tidak sesuai dengan harga voucher")
	}

	transaction := &Transaction{
		VoucherID:  voucherID,
		Quantity:   quantity,
		TotalPoint: totalPoint,
	}

	if err := CreateTransaction(transaction); err != nil {
		return nil, fmt.Errorf("gagal membuat transaksi: %w", err)
	}

	return transaction, nil
}

func GetDetailTransactionByID(id uint) (*Transaction, error) {
	transaction, err := GetTransactionByID(id)
	return transaction, err

}
