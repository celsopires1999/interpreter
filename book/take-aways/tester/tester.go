package tester

type Player interface {
	GetName() string
}

type Address struct {
	Street string
	City   string
	Zip    string
}

type Person struct {
	Name string
	Age  int
	Address
}

func (p Person) GetName() string {
	return p.Name
}

type Employee struct {
	Name string
	Age  int
	Address
}

func (e Employee) GetName() string {
	return e.Name
}

type Teams struct {
	Players []Player
}

func NewTeams() *Teams {
	return &Teams{}
}

func (t *Teams) AddPlayer(p Player) {
	t.Players = append(t.Players, p)
}
