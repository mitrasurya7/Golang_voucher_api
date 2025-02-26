package voucher

func CreateNewVoucher(brandID uint, name string, costInPoint int) (*Voucher, error) {
	voucher := &Voucher{
		BrandID:     brandID,
		Name:        name,
		CostInPoint: costInPoint,
	}
	err := CreateVoucher(voucher)
	return voucher, err
}
func GetSingleVoucherByID(id uint) (*Voucher, error) {
	voucher, err := GetVoucherByID(id)
	return voucher, err
}

func GetAllVoucherByBrandID(brandID uint) ([]Voucher, error) {
	vouchers, err := GetVoucherByBrandID(brandID)
	return vouchers, err
}
