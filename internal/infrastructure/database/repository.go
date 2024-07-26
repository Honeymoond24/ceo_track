package database

import "gorm.io/gorm"

type CeoRepositoryImpl struct{}

func (c CeoRepositoryImpl) CreateMany(db *gorm.DB, ceos []Ceo) {
	db.CreateInBatches(&ceos, 200)
}

func (c CeoRepositoryImpl) FindByBin(db *gorm.DB, bin string) Ceo {
	var ceo Ceo
	db.Where("company_bin = ?", bin).First(&ceo)
	return ceo
}

func (c CeoRepositoryImpl) FindByCompanyName(db *gorm.DB, companyName string) Ceo {
	var ceo Ceo
	db.Where("company_name = ?", companyName).First(&ceo)
	return ceo
}

func (c CeoRepositoryImpl) Update(db *gorm.DB, id uint, ceo Ceo) {
	db.Model(&Ceo{}).Where("id = ?", id).Updates(ceo)
}

type RegionRepositoryImpl struct{}

func (r RegionRepositoryImpl) CreateMany(db *gorm.DB, regionsDto []string) {
	var regions []Region
	for _, regionName := range regionsDto {
		regions = append(regions, Region{Name: regionName})
	}
	db.Create(&regions)
}

func (r RegionRepositoryImpl) FindByName(db *gorm.DB, name string) Region {
	var region Region
	db.Where("name = ?", name).First(&region)
	return region
}
