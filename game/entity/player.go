package entity

type Player struct {
	Entity
}

func NewPlayer() *Player {

	p := Player{}
	// #2 - here you can add your component
	// p.AddComponent(myComponent)

	return &p
}
