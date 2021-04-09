package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/kidussolo/gocrud/models"
)

func Wellcome(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Welcome to gocrud!"))
}

func NotifyNotFound(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("The route you are looking for is not found!"))
}

func AddItem(res http.ResponseWriter, req *http.Request) {
	b, _ := ioutil.ReadAll(req.Body)
	var item models.Item
	json.Unmarshal(b, &item)
	models.DB.Create(&item)
	res.Write([]byte("Item created Successfully"))
}

func GetItem(res http.ResponseWriter, req *http.Request) {
	var items []models.Item
	models.DB.Find(&items)
	fmt.Println(items)
	v, err := json.Marshal(items)
	if err != nil {
		panic(err)
	}
	res.Header().Add("Content-Type", "application/json")
	res.Write(v)
}
