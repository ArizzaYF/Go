package main

import (
    "fmt"
    "math/rand"
    "net/http"
    "time"
)

var opciones = []string{"piedra", "papel", "tijera"}

func elegirComputador() string {
    rand.Seed(time.Now().UnixNano())
    return opciones[rand.Intn(3)]
}

func decidirGanador(jugador, computador string) string {
    if jugador == computador { return "EMPATE" }
    gana := map[string]string{
        "piedra": "tijera",
        "tijera": "papel",
        "papel":  "piedra",
    }
    if gana[jugador] == computador { return "GANASTE" }
    return "PERDISTE"
}

func manejarInicio(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "static/index.html")
}

var svgIconos = map[string]string{
    "piedra": `<svg viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" width="36" height="36" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><path d="M3 20L8.5 9l4 6 2.5-4 6 9H3z"/></svg>`,
    "papel":  `<svg viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" width="36" height="36" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/><polyline points="14 2 14 8 20 8"/><line x1="16" y1="13" x2="8" y2="13"/><line x1="16" y1="17" x2="8" y2="17"/></svg>`,
    "tijera": `<svg viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" width="36" height="36" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><circle cx="6" cy="6" r="3"/><circle cx="6" cy="18" r="3"/><line x1="20" y1="4" x2="8.12" y2="15.88"/><line x1="14.47" y1="14.48" x2="20" y2="20"/><line x1="8.12" y1="8.12" x2="12" y2="12"/></svg>`,
}

func manejarJugar(w http.ResponseWriter, r *http.Request) {
    jugador := r.FormValue("jugador")
    if jugador != "piedra" && jugador != "papel" && jugador != "tijera" {
        fmt.Fprintf(w, "Opción inválida")
        return
    }

    computador := elegirComputador()
    resultado  := decidirGanador(jugador, computador)

    colorRes := "#4502FF"
    if resultado == "GANASTE" { colorRes = "#16A34A" }
    if resultado == "PERDISTE" { colorRes = "#DC2626" }

    fmt.Fprintf(w, `<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <title>Resultado</title>
    <link href="https://fonts.googleapis.com/css2?family=VT323&family=JetBrains+Mono&display=swap" rel="stylesheet">
    <style>
        *, *::before, *::after { box-sizing: border-box; margin: 0; padding: 0; }
        :root {
            --primary: #4502FF;
            --surface: #111114;
            --card:    #18181c;
            --border:  #2a2a30;
            --text:    #e8e8ed;
            --muted:   #52525e;
            --yellow:  #FFDA14;
        }
        body {
            font-family: 'VT323', monospace;
            background: var(--surface);
            color: var(--text);
            min-height: 100vh;
            display: grid;
            place-items: center;
        }
        .wrapper { display:flex; flex-direction:column; align-items:center; gap:32px; padding:32px; }
        .tag {
            font-family: 'JetBrains Mono', monospace;
            font-size: 12px; letter-spacing: 0.14em; text-transform: uppercase;
            color: var(--primary); border: 1px solid var(--primary); padding: 3px 12px;
        }
        h1 { font-family:'VT323',monospace; font-size:35px; letter-spacing:0.06em; text-transform:uppercase; line-height:1; margin-top:8px; }
        .card {
            background: var(--card); border: 1px solid var(--border);
            box-shadow: 4px 4px 0 var(--primary);
            padding: 32px 40px; display:flex; flex-direction:column; align-items:center; gap:24px; min-width:300px;
        }
        .versus { display:flex; align-items:center; gap:24px; justify-content:center; }
        .jugador { display:flex; flex-direction:column; align-items:center; gap:10px; }
        .icon-box {
            width:80px; height:80px; border:1px solid var(--border);
            background:var(--surface); display:grid; place-items:center;
            box-shadow:3px 3px 0 var(--primary); color:var(--text);
        }
        .icon-box svg { width:36px; height:36px; stroke:var(--text); fill:none; stroke-width:1.5; stroke-linecap:round; stroke-linejoin:round; }
        .jugador-label { font-family:'JetBrains Mono',monospace; font-size:11px; letter-spacing:0.12em; text-transform:uppercase; color:var(--muted); }
        .vs { font-family:'VT323',monospace; font-size:21px; color:var(--muted); }
        .divider { width:100%%; height:1px; background:var(--border); }
        .resultado { font-family:'VT323',monospace; font-size:35px; letter-spacing:0.08em; color:%s; }
        .btn {
            font-family:'JetBrains Mono',monospace; font-size:11px; letter-spacing:0.12em; text-transform:uppercase;
            padding:10px 24px; background:var(--primary); color:#fff;
            border:1px solid var(--primary); box-shadow:3px 3px 0 var(--text);
            text-decoration:none; display:inline-block;
            transition:box-shadow 80ms ease, transform 80ms ease;
        }
        .btn:hover { background:#3600cc; }
        .btn:active { box-shadow:0 0 0 var(--text); transform:translate(3px,3px); }
        .btn:focus-visible { outline:2px solid var(--yellow); outline-offset:3px; }
    </style>
</head>
<body>
    <div class="wrapper">
        <div style="display:flex;flex-direction:column;align-items:center;gap:6px;">
            <span class="tag">Resultado</span>
            <h1>Piedra · Papel · Tijera</h1>
        </div>
        <div class="card">
            <div class="versus">
                <div class="jugador">
                    <div class="icon-box">%s</div>
                    <span class="jugador-label">Tú</span>
                </div>
                <span class="vs">VS</span>
                <div class="jugador">
                    <div class="icon-box">%s</div>
                    <span class="jugador-label">CPU</span>
                </div>
            </div>
            <div class="divider"></div>
            <span class="resultado">%s</span>
            <a class="btn" href="/">Jugar de nuevo</a>
        </div>
    </div>
</body>
</html>`, colorRes, svgIconos[jugador], svgIconos[computador], resultado)
}

func main() {
    http.HandleFunc("/", manejarInicio)
    http.HandleFunc("/jugar", manejarJugar)
    fmt.Println("Servidor corriendo en http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
