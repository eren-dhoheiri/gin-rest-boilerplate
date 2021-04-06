package mobil

import "strconv"

type MobilFormatter struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Brand     string `json:"brand"`
	Type      string `json:"type"`
	Years     string `json:"years"`
	Condition string `json:"condition"`
	Price     string `json:"price"`
}

type BrandFormatter struct {
	ID        int    `json:"id"`
	Brand     string `json:"brand"`
}

func FormatBrand(brand BrandMobil) BrandFormatter {

	brandFormatter := BrandFormatter{}
	brandFormatter.ID = brand.ID
	brandFormatter.Brand = brand.Brand

	return brandFormatter

}

func FormatBrands(brands []BrandMobil) []BrandFormatter {
	
	brandsFormatter := []BrandFormatter{}

	for _, brand := range brands {
		brandFormatter := FormatBrand(brand)
		brandsFormatter = append(brandsFormatter, brandFormatter)
	}

	return brandsFormatter
}


func FormatMobil(mobil Mobil) MobilFormatter {

	mobilFormatter := MobilFormatter{}
	mobilFormatter.ID = mobil.ID
	mobilFormatter.Brand = mobil.MobilBrand.Brand
	mobilFormatter.Name = mobil.Name
	mobilFormatter.Condition = mobil.Condition
	mobilFormatter.Price = mobil.Price
	mobilFormatter.Type = mobil.Type
	mobilFormatter.Years = mobil.Years

	return mobilFormatter

}

func FormatMobilSave(mobil Mobil) MobilFormatter {

	mobilFormatter := MobilFormatter{}
	mobilFormatter.ID = mobil.ID
	mobilFormatter.Brand = strconv.Itoa(mobil.BrandID)
	mobilFormatter.Name = mobil.Name
	mobilFormatter.Condition = mobil.Condition
	mobilFormatter.Price = mobil.Price
	mobilFormatter.Type = mobil.Type
	mobilFormatter.Years = mobil.Years

	return mobilFormatter

}

func FormatMobils(mobils []Mobil) []MobilFormatter {

	mobilsFormatter := []MobilFormatter{}

	for _, mobil := range mobils {
		mobilFormatter := FormatMobil(mobil)
		mobilsFormatter = append(mobilsFormatter, mobilFormatter)
	}

	return mobilsFormatter
}