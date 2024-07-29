package data_source

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type Item struct {
	ItemId     int    `json:"itemId"`
	RegionName string `json:"name"`
	FileName   string
}

type ItemsResponse struct {
	Success bool   `json:"success"`
	List    []Item `json:"list"`
}

func (s *DataSource) getItemList(itemId int) []Item {
	url := fmt.Sprintf("https://stat.gov.kz/api/klazz/213/%d/ru", itemId)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	var response ItemsResponse
	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(&response); err != nil {
		fmt.Println(err)
	}
	fmt.Println("getItemList(", itemId, ") -> response:", response)
	return response.List
	//var itemIds []int
	//for _, item := range response.List { // _ is index, item is Item type
	//	itemIds = append(itemIds, item.ItemId)
	//}
	//return itemIds
}

type FileExportResponse struct {
	Success bool `json:"success"`
	Object  struct {
		Bucket   string `json:"bucket"`
		FileGuid string `json:"fileGuid"`
	} `json:"obj"`
}

func (s *DataSource) getDataFileLink(itemId int) string {
	url := fmt.Sprintf("https://stat.gov.kz/api/sbr/export/%d/ru", itemId)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	var response FileExportResponse
	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(&response); err != nil {
		fmt.Println(err)
	}
	dataFileLink := fmt.Sprintf(
		"https://stat.gov.kz/api/sbr/download?bucket=%s&guid=%s",
		response.Object.Bucket, response.Object.FileGuid,
	)
	//fmt.Println("dataFileLink:", dataFileLink)
	return dataFileLink
}

func (s *DataSource) fileDownload(url, fileName string) {
	resp, err := http.Get(url)

	dirPath := "files"
	filePath := filepath.Join(dirPath, fileName)
	// Create the directory if it doesn't exist
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}
	out, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(out)

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	fmt.Println("Downloaded:", filePath, "from", url)
}

func (s *DataSource) downloadData(itemId int) string {
	url := s.getDataFileLink(itemId)
	fileName := fmt.Sprintf("data_%d.xlsx", itemId)
	s.fileDownload(url, fileName)
	return fileName
}
