package brand

func CreateNewBrand(name string) (*Brand, error) {
	brand := &Brand{Name: name}
	err := CreateBrand(brand)
	return brand, err
}
