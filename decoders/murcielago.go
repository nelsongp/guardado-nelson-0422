package decoders

import "fmt"

type Murcielago struct{}

func (m *Murcielago) Notify(msg string) {
	codeMurcielago(msg)
}

func codeMurcielago(msg string) {
	fmt.Println("Estoy en murcielago: ", msg)
}
