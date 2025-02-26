package transaction

import "golang-voucher-api/pkg/database"

func CreateTransaction(transaction *Transaction) error {
	return database.DB.Create(transaction).Error
}
func GetTransactionByID(id uint) (*Transaction, error) {
	var transaction Transaction
	err := database.DB.Preload("Voucher").First(&transaction, id).Error

	return &transaction, err
}
