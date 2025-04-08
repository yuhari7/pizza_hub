package main

import (
	"net/http"

	"github.com/yuhari7/pizza_hub/delivery/http/chef"
	"github.com/yuhari7/pizza_hub/delivery/http/pizza"
	chefDomain "github.com/yuhari7/pizza_hub/domain/chef"
	pizzaDomain "github.com/yuhari7/pizza_hub/domain/pizza"
)

func main() {
	newChefChan := make(chan *chefDomain.Chef)
	chefRepo := chefDomain.NewChefRepository()
	newChefService := chefDomain.NewChefService(chefRepo, newChefChan)
	chefController := chef.NewChefController(newChefService)

	pizzaRepo := pizzaDomain.NewPizzaRepository()
	newPizzaService := pizzaDomain.NewPizzaService(pizzaRepo)
	pizzaController := pizza.NewPizzaController(newPizzaService)

	http.HandleFunc("/api/chef", chefController.RouteHandler)
	http.HandleFunc("/api/pizza", pizzaController.RouteHandler)
	http.ListenAndServe(":8080", nil)
}
