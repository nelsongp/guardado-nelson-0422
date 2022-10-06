package factory

//GetUserCodes sera la interfaz que utilizaremos para conocer los tipos de codificacion que utilizara el usuario
type GetUserCodes interface {
	GetCode() []string
}
