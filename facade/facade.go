package facade

import (
	"fmt"
	"strconv"

	"github.com/desarrolloitlm/guardado-nelson-0422/decoders"
	factory2 "github.com/desarrolloitlm/guardado-nelson-0422/factory"
	"github.com/desarrolloitlm/guardado-nelson-0422/message"
	"github.com/desarrolloitlm/guardado-nelson-0422/observer"
)

//Facade estructura para patron facade
type Facade struct {
	DecodeMorse      *decoders.Morse
	DecodeMd5        *decoders.Md5
	DecodeMurcielago *decoders.Murcielago
	DecodeSha12      *decoders.Sha12
}

//New nos permite crear una nueva facade
func New() Facade {
	return Facade{
		DecodeMorse:      &decoders.Morse{},
		DecodeMd5:        &decoders.Md5{},
		DecodeMurcielago: &decoders.Murcielago{},
		DecodeSha12:      &decoders.Sha12{},
	}
}

//GetEncryptedKeys sera el metodo que le vamos a exponer al usuario para que encripte
func (f *Facade) GetEncryptedKeys(user string, msg string) {
	fac2 := factory2.Factorys(user)
	if fac2 != nil {
		m := message.Message{}
		for l, cods := range fac2.GetCode() {
			imp := searchImpl(cods, f)
			m.AddObserver(strconv.Itoa(l), imp)
			m.Msg = msg
		}
		m.NotifyObservers()
	} else {
		fmt.Println("Invalid User/User not defined")
	}
}

//searchImpl nos servira para encontrar el tipo de implementacion a utilizar para subscribirlo al observador, en este caso igual lo mejor seria tenerlo en una base o configuraciones
func searchImpl(name string, f *Facade) observer.Observer {
	m := make(map[string]observer.Observer)
	m["morse"] = f.DecodeMorse
	m["md5"] = f.DecodeMd5
	m["sha12"] = f.DecodeSha12
	m["murcielago"] = f.DecodeMurcielago
	return m[name]
}
