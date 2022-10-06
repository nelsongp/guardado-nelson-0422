package observer

//Observer sera la interfaz donde vamos a definir el metodo a observar
type Observer interface {
	Notify(mesagge string)
}
