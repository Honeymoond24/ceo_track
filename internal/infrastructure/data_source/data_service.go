package data_source

import (
	"fmt"
)

type DataSource struct {
	DataFiles map[int]string
	Regions   map[int]string
}

func (s *DataSource) GetCeoData() []Item {
	baseItemId := s.getItemList(0)[0].ItemId
	items := s.getItemList(baseItemId)
	fmt.Println("getItemList(", baseItemId, ") -> items:", items)

	s.DataFiles = make(map[int]string)
	s.Regions = make(map[int]string)
	for i, item := range items {
		fmt.Println("item:", item)
		fileName := s.downloadData(item.ItemId)
		items[i].FileName = fileName
	}
	fmt.Println("items:", items)

	return items
}
