package main

import (
	"fmt"
	"strings"
)

func main() {
	// Mensaje a cifrar
	plaintext := "mensaje a cifrar"
	
	// Clave de cifrado (desplazamiento)
	key := 3
	
	// Cadena resultante del cifrado
	ciphertext := ""
	
	// Recorre cada caracter del mensaje a cifrar
	for _, ch := range plaintext {
		// Si el caracter es una letra mayúscula
		if ch >= 'A' && ch <= 'Z' {
			// Aplica el desplazamiento y obtiene el caracter cifrado
			ciphertext += string((ch-'A'+key)%26 + 'A')
		} else if ch >= 'a' && ch <= 'z' {
			// Si el caracter es una letra minúscula
			// Aplica el desplazamiento y obtiene el caracter cifrado
			ciphertext += string((ch-'a'+key)%26 + 'a')
		} else {
			// Si el caracter no es una letra, se agrega tal cual al resultado
			ciphertext += string(ch)
		}
	}
	
	// Imprime el mensaje cifrado
	fmt.Println(ciphertext)
}
