package middlewares

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-martini/martini"
)

func JSONParser(c martini.Context, res http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		fmt.Println("in the middle wwre")
		c.MapTo("JSONBody", b)

	}
	c.Next()
}
