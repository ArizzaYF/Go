package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    secreto  := rand.Intn(100) + 1
    intentos := 0
    maxIntentos := 7

    fmt.Println("¡Bienvenido! Adivina el número entre 1 y 100.")
    fmt.Printf("Tienes %d intentos.\n\n", maxIntentos)

    for {
        fmt.Printf("[%d/%d] Tu número: ", intentos+1, maxIntentos)
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

        if intentos == maxIntentos {
            fmt.Printf("Se acabaron los intentos. El número era: %d\n", secreto)
            break
        }
    }
}
