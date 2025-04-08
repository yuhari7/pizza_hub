package pizza_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/yuhari7/pizza_hub/domain/pizza"
	"github.com/yuhari7/pizza_hub/mocks"
)

func Test_pizzaImplementaion_GetAllMenu(t *testing.T) {
	type fields struct {
		repo pizza.Repository
	}

	tests := []struct {
		name   string
		fields fields
		want   []pizza.Pizza
	}{
		{
			name:   "Empty Menu",
			fields: fields{repo: &mocks.MockPizzaEmptyRepository{}},
			want:   nil,
		},
		{
			name:   "Non-empty Menu",
			fields: fields{repo: &mocks.MockPizzaRepository{}},
			want: []pizza.Pizza{
				{Name: "Cheese", Duration: 3 * time.Second},
				{Name: "BBQ", Duration: 5 * time.Second},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &pizza.PizzaImplementation{
				Repo: tt.fields.repo,
			}

			if got := p.GetAllMenu(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllMenu() = %v want %v", got, tt.want)
			}
		})
	}
}

func Test_pizzaImplementaion_GetMenuByKey(t *testing.T) {
	type fields struct {
		repo pizza.Repository
	}
	type args struct {
		req string
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *pizza.Pizza
	}{
		{
			name:   "Non-existing key",
			fields: fields{repo: &mocks.MockPizzaEmptyRepository{}},
			args:   args{req: "invalid_key"},
			want:   nil,
		},
		{
			name:   "Existing key",
			fields: fields{repo: &mocks.MockPizzaRepository{}},
			args:   args{req: "key1"},
			want:   &pizza.Pizza{Name: "Cheese", Duration: 3 * time.Second},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &pizza.PizzaImplementation{
				Repo: tt.fields.repo,
			}
			if got := p.GetMenuByKey(tt.args.req); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMenuByKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
