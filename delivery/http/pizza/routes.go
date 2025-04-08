package pizza

import (
	"net/http"

	pizzaDomain "github.com/yuhari7/pizza_hub/domain/pizza"
)

type controller struct {
	pizzaService pizzaDomain.Service
}

func NewPizzaController(pizzaService pizzaDomain.Service) *controller {
	return &controller{pizzaService: pizzaService}
}

func (c *controller) RouteHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		c.getListMenus(w, r)
	default:
		http.Error(w, "404 not found", http.StatusNotFound)
	}
}
