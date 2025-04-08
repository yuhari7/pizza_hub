package chef

type ChefImplementation struct {
	Repo        Repository
	ChefChannel chan *Chef
}

type Service interface {
	Add() (*Chef, int)
}

func NewChefService(repo Repository, chefChan chan *Chef) Service {
	return &ChefImplementation{
		Repo:        repo,
		ChefChannel: chefChan,
	}
}

func (c *ChefImplementation) Add() (*Chef, int) {
	newChef, total := c.Repo.Add()
	go func() {
		c.ChefChannel <- newChef
	}()

	return newChef, total
}
