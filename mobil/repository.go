package mobil

import "gorm.io/gorm"

type Repository interface {
	SaveMobil(mobil Mobil) (Mobil, error)
	FindAllMobil() ([]Mobil, error)
	FindByIDMobil(ID int) (Mobil, error)
	UpdateMobil(mobil Mobil) (Mobil, error)
	DeleteMobil(mobil Mobil) (Mobil, error)
	SaveBrand(brand BrandMobil) (BrandMobil, error)
	FindAllBrand() ([]BrandMobil, error)
	FindByIDBrand(ID int) (BrandMobil, error)
	DeleteBrand(brand BrandMobil) (BrandMobil, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) SaveBrand(brand BrandMobil) (BrandMobil, error) {
	err := r.db.Create(&brand).Error
	if err != nil {
		return brand, err
	}

	return brand, nil
}

func (r *repository) FindAllBrand() ([]BrandMobil, error) {
	var brands [] BrandMobil

	err := r.db.Find(&brands).Error
	if err != nil {
		return brands, err
	}

	return brands, nil
}

func (r *repository) FindByIDBrand(ID int) (BrandMobil, error) {
	var brand BrandMobil

	err := r.db.Where("ID = ?", ID).Find(&brand).Error
	if err != nil {
		return brand, err
	}

	return brand, nil
}


func (r *repository) DeleteBrand(brand BrandMobil) (BrandMobil, error){
	err := r.db.Select("id").Delete(&brand).Error
	if err != nil {
		return brand, err
	}

	return brand, nil
}



func (r *repository) SaveMobil(mobil Mobil) (Mobil, error) {
	err := r.db.Create(&mobil).Error
	if err != nil {
		return mobil, err
	}

	return mobil, nil
}

func (r *repository) FindAllMobil() ([]Mobil, error) {
	var mobils [] Mobil

	err := r.db.Preload("MobilBrand", "ID != 0").Find(&mobils).Error
	if err != nil {
		return mobils, err
	}

	return mobils, nil
}

func (r *repository) FindByIDMobil(ID int) (Mobil, error) {
	var mobil Mobil

	err := r.db.Where("ID = ?", ID).Preload("MobilBrand", "ID != 0").Find(&mobil).Error
	if err != nil {
		return mobil, err
	}

	return mobil, nil
}


func (r *repository) UpdateMobil(mobil Mobil) (Mobil, error){
	err := r.db.Save(&mobil).Error
	if err != nil {
		return mobil, err
	}

	return mobil, nil
}

func (r *repository) DeleteMobil(mobil Mobil) (Mobil, error){
	err := r.db.Select("id").Delete(&mobil).Error
	if err != nil {
		return mobil, err
	}

	return mobil, nil
}