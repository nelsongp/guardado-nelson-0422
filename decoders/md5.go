package decoders

import "fmt"

//Md5 estructura de md5
type Md5 struct{}

/*codeMorse funcion que recibe un string y lo codifica a Morse*/
func (m *Md5) Notify(msg string) {
	codeMd5(msg)
}

//codeMd5 funcion que nos codificara a formato md5
func codeMd5(msg string) {
	fmt.Println("Estoy en md5: ", msg)
}
