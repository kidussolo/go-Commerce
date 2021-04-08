package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Wellcome(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Welcome to gocrud!"))
}

func NotifyNotFound(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("The route you are looking for is not found!"))
}

func AddItem(res http.ResponseWriter, req *http.Request) {
	b, _ := ioutil.ReadAll(req.Body)

	fmt.Println(b)
}
