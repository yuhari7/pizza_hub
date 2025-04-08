package mocks

import (
	"time"

	"github.com/yuhari7/pizza_hub/domain/pizza"
)

type MockPizzaRepository struct{}
type MockPizzaEmptyRepository struct{}

func (m *MockPizzaRepository) GetMenuByKey(key string) *pizza.Pizza {
	return &pizza.Pizza{Name: "Cheese", Duration: 3 * time.Second}
}

func (m *MockPizzaRepository) GetAllMenu() []pizza.Pizza {
	return []pizza.Pizza{
		{Name: "Cheese", Duration: 3 * time.Second},
		{Name: "BBQ", Duration: 5 * time.Second},
	}
}

func (m *MockPizzaEmptyRepository) GetMenuByKey(key string) *pizza.Pizza {
	return nil
}

func (m *MockPizzaEmptyRepository) GetAllMenu() []pizza.Pizza {
	return nil
}
