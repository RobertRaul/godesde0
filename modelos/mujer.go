package modelos

type Mujer struct {
	Hombre
	EsMadre bool
}

func (m *Mujer) Respirar()    { m.respirando = true }
func (m *Mujer) Pensar()      { m.pensar = true }
func (m *Mujer) Comer()       { m.comiendo = "Si comer" }
func (m *Mujer) Sexo() string { return "Mujer" }
func (m *Mujer) EstaLive()    { m.vivo = true }
