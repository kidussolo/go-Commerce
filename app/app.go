package app

import (
	"github.com/go-martini/martini"
	"github.com/kidussolo/gocrud/app/controllers"
)

func Start() {
	port := ":8082"
	baseUrl := "/go-commerce"

	// Initialize martini server
	m := martini.Classic()

	// Add route and controllers
	m.Post(baseUrl+"/items", controllers.AddItem)
	m.Get(baseUrl+"/welcome", controllers.Wellcome)
	m.NotFound(controllers.NotifyNotFound)

	// Run server on port
	m.RunOnAddr(port)

}
