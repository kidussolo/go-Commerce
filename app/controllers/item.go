package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func AddItem(res http.ResponseWriter, req *http.Request) {
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(res, err.Error(), 500)
	}

	fmt.Println(b)
}
