package main

import (
	"net/http"

	"github.com/yuhari7/pizza_hub/delivery/http/chef"
	chefDomain "github.com/yuhari7/pizza_hub/domain/chef"
)

func main() {
	newChefChan := make(chan *chefDomain.Chef)
	chefRepo := chefDomain.NewChefRepository()
	newChefService := chefDomain.NewChefService(chefRepo, newChefChan)
	chefController := chef.NewChefController(newChefService)

	http.HandleFunc("/api/chef", chefController.RouteHandler)
	http.ListenAndServe(":8080", nil)
}
