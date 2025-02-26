package voucher

import "golang-voucher-api/pkg/database"

func CreateVoucher(voucher *Voucher) error {
	return database.DB.Create(voucher).Error
}

func GetVoucherByID(id uint) (*Voucher, error) {
	var voucher Voucher
	err := database.DB.First(&voucher, id).Error
	return &voucher, err
}

func GetVoucherByBrandID(brandID uint) ([]Voucher, error) {
	var vouchers []Voucher
	err := database.DB.Where("brand_id = ?", brandID).Find(&vouchers).Error
	return vouchers, err
}
