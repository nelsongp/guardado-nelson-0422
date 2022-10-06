package observer

/*Observable interfaz que lleva el manejo de los observadores*/
type Observable interface {
	AddObserver(name string, o Observer)
	RemoveObserver(name string)
	NotifyObserver()
}
