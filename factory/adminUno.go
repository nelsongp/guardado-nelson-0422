package factory

//AdminUno nos servira para definir los tipos de codificacion que utilizara el admin uno
type AdminUno struct {
	listCodes []string
}

//GetCode interfaz que vamos a utilizar para exponer en el factory
func (a *AdminUno) GetCode() []string {
	//lo ideal fuera sacarlo de una base de datos o de configuraciones
	a.listCodes = []string{"morse", "md5"}
	return a.listCodes
}
