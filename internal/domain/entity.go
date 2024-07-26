package domain

type Ceo struct {
	Id          uint
	CompanyBin  string
	CompanyName string
	CeoFullName string
	RegionID    uint
}

type Region struct {
	Id   uint
	Name string
}
