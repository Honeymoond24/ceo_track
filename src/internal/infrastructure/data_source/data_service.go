package data_source

import (
	"fmt"
)

type DataSource struct {
	DataFiles map[int]string
}

func (s *DataSource) GetCeoData() map[int]string {
	baseItemId := s.getItemList(0)[0]
	itemIds := s.getItemList(baseItemId)
	fmt.Println("getItemList(", baseItemId, ") -> itemIds:", itemIds)

	s.DataFiles = make(map[int]string)
	for _, itemId := range itemIds {
		fmt.Println("itemId:", itemId)
		fileName := s.downloadData(itemId)
		s.DataFiles[itemId] = fileName
		break
	}
	fmt.Println("dataFiles:", s.DataFiles)
	return s.DataFiles
}
