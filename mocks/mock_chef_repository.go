package mocks

import "github.com/yuhari7/pizza_hub/domain/chef"

type MockChefRepository struct {
	chefChannel chan *chef.Chef
}

func (m *MockChefRepository) Add() (*chef.Chef, int) {
	return &chef.Chef{ID: 1}, 1
}

func (m *MockChefRepository) ListChefs() []*chef.Chef {
	return []*chef.Chef{
		{ID: 1},
		{ID: 2},
	}
}

func (m *MockChefRepository) GetChefsCount() int {
	return 2
}

type MockChefEmptyRepository struct{}

func (m *MockChefEmptyRepository) Add() (*chef.Chef, int) {
	return nil, 0
}

func (m *MockChefEmptyRepository) ListChefs() []*chef.Chef {
	return nil
}

func (m *MockChefEmptyRepository) GetChefsCount() int {
	return 0
}
