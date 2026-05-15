package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    secreto := rand.Intn(100) + 1
    intentos := 0

    fmt.Println("¡Bienvenido! Adivina el número entre 1 y 100.")

    for {
        fmt.Print("Tu número: ")
        var intento int
        fmt.Scan(&intento)
        intentos++

        if intento < secreto {
            fmt.Println("¡Más alto!")
        } else if intento > secreto {
            fmt.Println("¡Más bajo!")
        } else {
            fmt.Printf("¡Correcto en %d intentos!\n", intentos)
            break
        }
    }
}
