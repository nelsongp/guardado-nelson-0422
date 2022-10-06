package message

import "github.com/desarrolloitlm/guardado-nelson-0422/observer"

//Message estructura que va a monitorear los observadores a agregar y el mensaje a descifrar
type Message struct {
	observers map[string]observer.Observer
	Msg       string
}

//AddObserver implementacion de la interfaz observer
func (m *Message) AddObserver(name string, o observer.Observer) {
	if m.observers == nil {
		m.observers = make(map[string]observer.Observer)
	}
	m.observers[name] = o
}

//RemoveObserver implementacion de la interfaz observer
func (m *Message) RemoveObserver(name string) {
	delete(m.observers, name)
}

//NotifyObservers implementacion de la interfaz observer
func (m *Message) NotifyObservers() {
	for _, v := range m.observers {
		v.Notify(m.Msg)
	}
}
