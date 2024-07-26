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
	//items = items[:2]
	s.DataFiles = make(map[int]string)
	s.Regions = make(map[int]string)
	//tmpCounter := 0
	for i, item := range items {
		//if tmpCounter++; tmpCounter > 2 {
		//	continue
		//}
		fmt.Println("item:", item)
		fileName := s.downloadData(item.ItemId)
		items[i].FileName = fileName
		//item.FileName = fileName
		//fmt.Println("fileName:", item.FileName, 'a', fileName)
	}
	fmt.Println("items:", items)

	return items
}
