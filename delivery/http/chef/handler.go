package chef

import (
	"encoding/json"
	"net/http"

	"github.com/yuhari7/pizza_hub/common"
	"github.com/yuhari7/pizza_hub/delivery/http/chef/model"
)

func (c *controller) addChef(w http.ResponseWriter, r *http.Request) {
	chef, total := c.chefService.Add()
	response := model.AddChefResponse{
		ID:         chef.ID,
		TotalChefs: total,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(common.SuccessResponseWithData(response, "success"))
}
