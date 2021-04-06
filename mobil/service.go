package mobil

type Service interface {
	GetMobil() ([]Mobil, error)
	GetMobilByID(input GetMobilDetailInput) (Mobil, error)
	CreateMobil(input CreateMobilInput) (Mobil, error)
	UpdateMobil(inputID GetMobilDetailInput, inputData CreateMobilInput) (Mobil, error)
	DeleteMobil(input GetMobilDetailInput) (Mobil, error)
	GetBrand() ([]BrandMobil, error)
	GetBrandByID(input GetBrandDetailInput) (BrandMobil, error)
	CreateBrand(input CreateBrandInput) (BrandMobil, error)
	DeleteBrand(input GetBrandDetailInput) (BrandMobil, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetBrand() ([]BrandMobil, error) {

	brands, err := s.repository.FindAllBrand()
	if err != nil {
		return brands, err
	}

	return brands, nil

}

func (s *service) GetBrandByID(input GetBrandDetailInput) (BrandMobil, error) {

	brand, err := s.repository.FindByIDBrand(input.ID)

	if err != nil {
		return brand, err
	}

	return brand, nil

}

func (s *service) CreateBrand(input CreateBrandInput) (BrandMobil, error) {

	brand := BrandMobil{}
	brand.Brand = input.Brand

	newBrand, err := s.repository.SaveBrand(brand)
	if err != nil {
		return newBrand, err
	}

	return newBrand, nil
}

func (s *service) DeleteBrand(input GetBrandDetailInput) (BrandMobil, error) {
	brand, err := s.repository.FindByIDBrand(input.ID)

	if err != nil {
		return brand, err
	}

	deletedBrand, err := s.repository.DeleteBrand(brand)
	if err != nil {
		return deletedBrand, err
	}

	return deletedBrand, nil
}

func (s *service) GetMobil() ([]Mobil, error) {

	mobils, err := s.repository.FindAllMobil()
	if err != nil {
		return mobils, err
	}

	return mobils, nil

}

func (s *service) GetMobilByID(input GetMobilDetailInput) (Mobil, error) {

	mobil, err := s.repository.FindByIDMobil(input.ID)

	if err != nil {
		return mobil, err
	}

	return mobil, nil

}

func (s *service) CreateMobil(input CreateMobilInput) (Mobil, error) {

	mobil := Mobil{}
	mobil.BrandID = input.Brand
	mobil.Type = input.Type
	mobil.Years = input.Years
	mobil.Name = input.Name
	mobil.Price = input.Price
	mobil.Condition = input.Condition

	newMobil, err := s.repository.SaveMobil(mobil)
	if err != nil {
		return newMobil, err
	}

	return newMobil, nil
}

func (s *service) UpdateMobil(inputID GetMobilDetailInput, inputData CreateMobilInput) (Mobil, error) {

	mobil, err := s.repository.FindByIDMobil(inputID.ID)
	if err != nil {
		return mobil, err
	}

	mobil.BrandID = inputData.Brand
	mobil.Type = inputData.Type
	mobil.Years = inputData.Years
	mobil.Name = inputData.Name
	mobil.Price = inputData.Price
	mobil.Condition = inputData.Condition

	updatedMobil, err := s.repository.UpdateMobil(mobil)
	if err != nil {
		return updatedMobil, err
	}

	return updatedMobil, nil
}

func (s *service) DeleteMobil(input GetMobilDetailInput) (Mobil, error) {
	mobil, err := s.repository.FindByIDMobil(input.ID)

	if err != nil {
		return mobil, err
	}

	deletedMobil, err := s.repository.DeleteMobil(mobil)
	if err != nil {
		return deletedMobil, err
	}

	return deletedMobil, nil
}