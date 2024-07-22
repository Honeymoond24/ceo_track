package application

type IDataSource interface {
	GetCeoData() map[int]string
	downloadData(itemId int) string
	fileDownload(url, fileName string)
	getDataFileLink(itemId int) string
	getItemList(itemId int) []int
}
