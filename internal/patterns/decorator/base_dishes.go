package decorator

type Pizza struct {
	Size string
}

func NewPizza(size string) *Pizza {
	return &Pizza{Size: size}
}

func (p *Pizza) GetDescription() string {
	return p.Size + " пицца"
}

func (p *Pizza) GetPrice() float64 {
	switch p.Size {
	case "маленькая":
		return 300.0
	case "средняя":
		return 450.0
	case "большая":
		return 600.0
	default:
		return 450.0
	}
}

type Pasta struct {
	Type string
}

func NewPasta(pastaType string) *Pasta {
	return &Pasta{Type: pastaType}
}

func (p *Pasta) GetDescription() string {
	return "паста " + p.Type
}

func (p *Pasta) GetPrice() float64 {
	switch p.Type {
	case "карбонара":
		return 350.0
	case "болоньезе":
		return 380.0
	default:
		return 350.0
	}
}

type Salad struct {
	Name string
}

func NewSalad(name string) *Salad {
	return &Salad{Name: name}
}

func (s *Salad) GetDescription() string {
	return "салат " + s.Name
}

func (s *Salad) GetPrice() float64 {
	switch s.Name {
	case "цезарь":
		return 280.0
	case "греческий":
		return 250.0
	default:
		return 250.0
	}
}
