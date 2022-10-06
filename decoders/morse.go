package decoders

import (
	"fmt"
	"strings"
)

//Morse estructura para morse
type Morse struct{}

//Notify es el notificador del observer
func (m *Morse) Notify(msg string) {
	codeMorse(msg)
}

/*codeMorse funcion que recibe un string y lo codifica a Morse*/
func codeMorse(msg string) {
	n := strings.ToLower(msg)
	var s []string
	for _, l := range n {
		s = append(s, searchLetter(string(l)))
	}
	fmt.Println("Estoy en morse: ", strings.Join(s[:], ""))
}

//searchLetter nos ayudara a buscar el tipo de letra para codificarla a morse
func searchLetter(letra string) string {
	m := make(map[string]string)
	m["a"] = ".-"
	m["b"] = "-..."
	m["c"] = ".-.-"
	m["d"] = "-.."
	m["e"] = "."
	m["f"] = "..-."
	m["g"] = "--."
	m["h"] = "...."
	m["i"] = ".."
	m["j"] = ".---"
	m["k"] = "-.-"
	m["l"] = "-.-"
	m["m"] = "--"
	m["n"] = "-.."
	m["ñ"] = "--.--"
	m["o"] = "---"
	m["p"] = ".--."
	m["q"] = "--.-"
	m["r"] = ".-."
	m["s"] = "..."
	m["t"] = "-"
	m["u"] = "..-"
	m["v"] = "...-"
	m["w"] = ".--"
	m["x"] = "-..-"
	m["y"] = "-.--"
	m["z"] = "--.."

	return m[letra]
}
