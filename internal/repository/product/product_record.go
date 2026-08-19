package product

type ProductRecord struct {
	ID          string   `gorm:"primaryKey"`
	Name        string   `gorm:"not null"`
	Description *string  `gorm:"type:text"`
	SalePrice   *float64 `gorm:"type:decimal(12,2)"`
	Price       float64  `gorm:"not null;type:decimal(12,2)"`
}

func (ProductRecord) TableName() string {
	return "products"
}
