package service

type Greeter struct{}

func NewGreeter() *Greeter {
	return &Greeter{}
}

func (g *Greeter) Greet() string {
	return "Hello, world!"
}
