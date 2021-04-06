package mobil

import "time"

type Mobil struct {
	ID             int
	MobilBrand	   BrandMobil `gorm:"foreignKey:BrandID"`
	BrandID		   int
	Type    	   string
	Years          string
	Name           string
	Condition      string
	Price		   string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}


type BrandMobil struct {
	ID             int 
	Brand          string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}