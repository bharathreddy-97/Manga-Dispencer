package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/bharathreddy-97/Manga-Dispencer/constants"
	"github.com/bharathreddy-97/Manga-Dispencer/models"
)

func HandleGetAllBooksList(rw http.ResponseWriter, r *http.Request) {
	byteData, readError := os.ReadFile(constants.BooksJsonPath)
	if readError != nil {
		rw.Write([]byte(readError.Error()))
		return
	}
	var arrayData = []interface{}{}
	unmarshallError := json.Unmarshal(byteData, &arrayData)
	if unmarshallError != nil {
		rw.Write([]byte(unmarshallError.Error()))
		return
	}
	var books = []models.Book{}
	for _, arrayObject := range arrayData {
		if dictData, ok := arrayObject.(map[string]interface{}); ok {
			var book = models.Book{}
			book.GetDataFromDict(dictData)
			books = append(books, book)
		}
	}

	responseData, marshallError := json.Marshal(&books)
	if marshallError != nil {
		rw.Write([]byte(marshallError.Error()))
		return
	}
	rw.Write(responseData)
}
