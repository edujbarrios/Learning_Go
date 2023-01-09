package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

func main() {
	ctx := context.Background()

	// Obtén la clave de API de Google Cloud Platform.
	apiKey := "YOUR_API_KEY"

	client, err := translate.NewClient(ctx, translate.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Define el idioma de origen y el idioma de destino.
	sourceLanguage := language.English
	targetLanguage, err := language.Parse("es")
	if err != nil {
		log.Fatalf("Failed to parse target language: %v", err)
	}

	// Realiza la traducción.
	translation, err := client.Translate(ctx, []string{"Hello"}, targetLanguage, &translate.Options{
		Source: sourceLanguage,
	})
	if err != nil {
		log.Fatalf("Failed to translate text: %v", err)
	}

	// Imprime el resultado.
	fmt.Printf("Text: %v\n", translation[0].Text)
}
