package brand

import "golang-voucher-api/pkg/database"

func CreateBrand(brand *Brand) error {
	return database.DB.Create(brand).Error
}

func GetBrandByID(id uint) (*Brand, error) {
	var brand Brand
	err := database.DB.First(&brand, id).Error
	return &brand, err
}
