package main

import (
    "fmt"
    "math/rand"
    "time"
)

// Struct del personaje: agrupa todos sus datos en un solo tipo
type Personaje struct {
    Nombre  string
    Vida    int
    Ataque  int
    Defensa int
}

// Struct de habilidad: nombre y cuánto bonus de ataque da
type Habilidad struct {
    Nombre string
    Poder  int
}

func main() {
    rand.Seed(time.Now().UnixNano())

    // Creamos los dos personajes con sus estadísticas iniciales
    heroe    := Personaje{"Héroe", 100, 20, 5}
    monstruo := Personaje{"Dragón", 80, 15, 3}
    turno    := 1

    // Habilidades que el JUGADOR puede elegir cada turno
    var habilidadesHeroe = []Habilidad{
        {"Golpe normal", 0},
        {"Golpe fuerte", 10},
        {"Magia de fuego", 20},
    }

    // Habilidades que el DRAGÓN usa al azar
    var habilidadesDragon = []Habilidad{
        {"Zarpazo",       5},
        {"Llamarada",     15},
        {"Rugido Mortal", 25},
    }

    // El bucle corre mientras los dos estén vivos
    for heroe.EstaVivo() && monstruo.EstaVivo() {
        fmt.Printf("\n========== TURNO %d ==========\n", turno)
        fmt.Printf("Tu vida: %d  |  Vida del Dragón: %d\n\n", heroe.Vida, monstruo.Vida)

        // --- Turno del jugador ---
        hab := elegirHabilidad(habilidadesHeroe)
        heroe.Ataque += hab.Poder          // Bonus temporal
        heroe.Atacar(&monstruo)
        heroe.Ataque -= hab.Poder          // Quitamos el bonus para el próximo turno

        // --- Turno del dragón (solo si sigue vivo) ---
        if monstruo.EstaVivo() {
            habDragon := habilidadesDragon[rand.Intn(len(habilidadesDragon))]
            fmt.Printf("⚠️  El Dragón prepara: %s (poder: +%d)\n\n", habDragon.Nombre, habDragon.Poder)
            monstruo.Ataque += habDragon.Poder
            monstruo.Atacar(&heroe)
            monstruo.Ataque -= habDragon.Poder
        }
        turno++
    }

    // --- Resultado final ---
    fmt.Println("\n==============================")
    if heroe.EstaVivo() {
        fmt.Println("¡Ganaste! El héroe venció al dragón. 🏆")
    } else {
        fmt.Println("El dragón ganó. Inténtalo de nuevo. 💀")
    }
}

// EstaVivo devuelve true si el personaje tiene más de 0 de vida
func (p *Personaje) EstaVivo() bool {
    return p.Vida > 0
}

// Atacar calcula el daño, aplica variación aleatoria y resta la vida al objetivo.
// Usamos *Personaje (puntero) para modificar el original, no una copia.
func (p *Personaje) Atacar(objetivo *Personaje) {
    dano := p.Ataque - objetivo.Defensa
    if dano < 1 {
        dano = 1
    }
    // Variación aleatoria ±5
    variacion := rand.Intn(11) - 5
    dano += variacion
    if dano < 1 {
        dano = 1
    }
    objetivo.Vida -= dano
    fmt.Printf("%s ataca a %s por %d de daño!\n", p.Nombre, objetivo.Nombre, dano)
    fmt.Printf("%s tiene %d de vida restante.\n\n", objetivo.Nombre, objetivo.Vida)
}

// elegirHabilidad muestra el menú y devuelve la habilidad que eligió el jugador.
func elegirHabilidad(habilidades []Habilidad) Habilidad {
    fmt.Println("Elige tu habilidad:")
    for i, h := range habilidades {
        fmt.Printf("  %d. %s (poder: +%d)\n", i+1, h.Nombre, h.Poder)
    }
    var opcion int
    fmt.Scan(&opcion)
    if opcion < 1 || opcion > len(habilidades) {
        fmt.Println("Opción inválida, se usó:", habilidades[0].Nombre)
        return habilidades[0]
    }
    return habilidades[opcion-1]
}
