package decoders

import "fmt"

//Sha12 estructura para morse
type Sha12 struct{}

//Notify es el notificador del observer
func (m *Sha12) Notify(msg string) {
	codeSha12(msg)
}

//codeSha12 funcion que nos codificara a formato sha12
func codeSha12(msg string) {
	fmt.Println("Estoy en sha12: ", msg)
}
