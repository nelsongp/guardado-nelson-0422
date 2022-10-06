package factory

//UserUno nos servira para definir los tipos de codificacion que utilizara el user uno
type UserUno struct {
	listCodes []string
}

//GetCode interfaz que vamos a utilizar para exponer en el factory
func (u *UserUno) GetCode() []string {
	//lo ideal fuera sacarlo de una base de datos o de configuraciones
	u.listCodes = []string{"sha12", "murcielago"}
	return u.listCodes
}
