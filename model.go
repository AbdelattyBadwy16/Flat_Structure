package main

type Brand struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"foreignKey:BrandID"`
}

type Product struct {
	ID      uint   `gorm:"primaryKey"`
	Name    string `json:"name"`
	BrandID uint   `json:"brand_id"`
	Brand   Brand
	Tags    []Tag `gorm:"many2many:product_tags;"`
}

type Tag struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:product_tags;"`
}
