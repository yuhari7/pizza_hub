package pizza

type PizzaImplementation struct {
	Repo Repository
}

type Service interface {
	GetMenuByKey(req string) *Pizza
	GetAllMenu() []Pizza
}

func NewPizzaService(repo Repository) Service {
	return &PizzaImplementation{
		Repo: repo,
	}
}

func (p *PizzaImplementation) GetMenuByKey(req string) *Pizza {
	return p.Repo.GetMenuByKey(req)
}

func (p *PizzaImplementation) GetAllMenu() []Pizza {
	return p.Repo.GetAllMenu()
}
