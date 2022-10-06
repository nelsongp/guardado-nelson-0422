package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/desarrolloitlm/guardado-nelson-0422/facade"
)

func main() {
	u := readUser()
	m := readMessage()
	f := facade.New()
	f.GetEncryptedKeys(u, m)
}

//readUser funcion que nos ayudara a capturar el nombre del usuario
func readUser() string {
	fmt.Println("Digite el nombre del usuario: ")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\r')
	if err != nil {
		log.Fatalf("no se pudo leer lo que digitó el usuario: %v", err)
	}
	text = strings.TrimSuffix(text, "\r")
	return text
}

//readMessage funcion que nos ayudara a capturar el nombre del usuario
func readMessage() string {
	fmt.Println("Digite el texto a encriptar: ")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\r')
	if err != nil {
		log.Fatalf("no se pudo leer lo que digitó el usuario: %v", err)
	}
	text = strings.TrimSuffix(text, "\r")
	return text
}
