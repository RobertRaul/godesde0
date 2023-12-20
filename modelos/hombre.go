package modelos

type Hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensar     bool
	vivo       bool
	comiendo   string
}

func (h *Hombre) Respirar()    { h.respirando = true }
func (h *Hombre) Pensar()      { h.pensar = true }
func (h *Hombre) Comer()       { h.comiendo = "Si comer" }
func (h *Hombre) Sexo() string { return "Hombre" }
func (h *Hombre) EstaLive()    { h.vivo = true }
