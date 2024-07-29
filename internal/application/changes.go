package application

type CeoNew struct {
	CompanyBin     string
	NewCeoFullName string
}
type CeoChange struct {
	CeoNew
	OldCeoFullName string
}
type CeoChanges struct {
	RegionID uint
	New      []CeoNew
	Changes  []CeoChange
}

func (c *CeoChanges) MakeMessageReport() string {
	var message string
	if len(c.New) > 0 {
		message += "<b>Новые записи:</b>\n"
		for _, ceo := range c.New {
			message += ceo.CompanyBin + " - " + ceo.NewCeoFullName + "\n"
		}
	}
	if len(c.Changes) > 0 {
		message += "<b>Смены директоров:</b>\n"
		for _, ceo := range c.Changes {
			message += ceo.CompanyBin + " - " + ceo.NewCeoFullName + " -> " + ceo.OldCeoFullName + "\n"
		}
	}
	if len(c.New) == 0 && len(c.Changes) == 0 {
		message = "Нет изменений"
	}
	return message
}
