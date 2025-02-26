package voucher

type Voucher struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	BrandID     uint   `gorm:"not null" json:"brand_id"`
	Name        string `gorm:"not null" json:"name"`
	CostInPoint int    `gorm:"not null" json:"cost_in_point"`
}
