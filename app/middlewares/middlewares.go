package middlewares

import (
	"io/ioutil"
	"net/http"

	"github.com/go-martini/martini"
)

func JSONParser(c martini.Context, res http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(res, err.Error(), 500)
		}

		c.MapTo("JSONBODY", b)

	}
	c.Next()
}
