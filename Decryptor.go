package main

import (
	"fmt"
	"strings"
)

func main() {
	// Mensaje cifrado
	ciphertext := "vhw th qrls" //es un de prueba sin traduccion 
	
	// Clave de cifrado (desplazamiento)
	key := 3
	
	// Cadena resultante del descifrado
	plaintext := ""
	
	// Recorre cada caracter del mensaje cifrado
	for _, ch := range ciphertext {
		// Si el caracter es una letra mayúscula
		if ch >= 'A' && ch <= 'Z' {
			// Aplica el desplazamiento y obtiene el caracter descifrado
			plaintext += string((ch-'A'-key+26)%26 + 'A')
		} else if ch >= 'a' && ch <= 'z' {
			// Si el caracter es una letra minúscula
			// Aplica el desplazamiento y obtiene el caracter descifrado
			plaintext += string((ch-'a'-key+26)%26 + 'a')
		} else {
			// Si el caracter no es una letra, se agrega tal cual al resultado
			plaintext += string(ch)
		}
	}
	
	// Imprime el mensaje descifrado
	fmt.Println(plaintext)
}
