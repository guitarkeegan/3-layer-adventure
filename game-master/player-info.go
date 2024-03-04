package gamemaster

type Player struct {
	Name  string
	Loves string
	Fears string
}

func NewPlayer(name, loves, fears string) Player {
	return Player{
		Name:  name,
		Loves: loves,
		Fears: fears,
	}
}
