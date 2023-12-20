package interfaces

type animal interface {
	respirar()
	comer()
	esCarnivoro() bool
	EstaLive() bool
}
