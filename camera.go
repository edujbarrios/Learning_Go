package main

import (
    "fmt"
    "image"
    _ "image/jpeg"
    "log"
    "os"

    "github.com/blackjack/webcam"
)

func main() {
    // Abrir la cámara
    cam, err := webcam.Open("/dev/video0")
    if err != nil {
        log.Fatal(err)
    }
    defer cam.Close()

    // Configurar la cámara
    format := webcam.PixelFormat(webcam.PixelFormatYUYV)
    if err = cam.SetImageFormat(640, 480, format); err != nil {
        log.Fatal(err)
    }
    if err = cam.StartStreaming(); err != nil {
        log.Fatal(err)
    }
    defer cam.StopStreaming()

    // Tomar una imagen
    frame, _, err := cam.GetFrame()
    if err != nil {
        log.Fatal(err)
    }

    // Decodificar el frame como una imagen
    img, _, err := image.Decode(frame)
    if err != nil {
        log.Fatal(err)
    }

    // Guardar la imagen a un archivo
    file, err := os.Create("image.jpg")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    if err = jpeg.Encode(file, img, &jpeg.Options{Quality: jpeg.DefaultQuality}); err != nil {
        log.Fatal(err)
    }
    fmt.Println("Imagen guardada en image.jpg")
}
