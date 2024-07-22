package application

type CeoRepository interface {
	Create(ceo CeoDTO) int
	Get(ceoId int) CeoDTO
	Update(ceo CeoDTO)
	Delete(ceoId int)
}
