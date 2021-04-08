package controllers

import (
	"encoding/json"
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
