package database

type Ceo struct {
	ID          uint   `gorm:"primarykey"`
	CompanyBin  string `gorm:"index:idx_company_bin"`
	CompanyName string `gorm:"index:idx_company_name"`
	FullName    string
	RegionID    uint
}

type Region struct {
	ID   uint `gorm:"primarykey"`
	Name string
}
