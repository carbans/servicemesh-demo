package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {
    // Configurar el endpoint de destino utilizando una variable de entorno
    targetEndpoint := os.Getenv("TARGET_ENDPOINT")
    if targetEndpoint == "" {
        log.Fatal("La variable de entorno TARGET_ENDPOINT no está configurada")
    }

    // Manejador para la ruta raíz
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // Realizar la solicitud al endpoint de destino
        resp, err := http.Get(targetEndpoint)
        if err != nil {
            http.Error(w, fmt.Sprintf("Error al realizar la solicitud al endpoint de destino: %v", err), http.StatusInternalServerError)
            return
        }
        defer resp.Body.Close()

        // Leer la respuesta del endpoint de destino y enviarla como respuesta al cliente
        body := make([]byte, 0)
        _, err = resp.Body.Read(body)
        if err != nil {
            http.Error(w, fmt.Sprintf("Error al leer la respuesta del endpoint de destino: %v", err), http.StatusInternalServerError)
            return
        }

        // Escribir la respuesta del endpoint de destino como respuesta al cliente
        w.WriteHeader(resp.StatusCode)
        _, err = w.Write(body)
        if err != nil {
            http.Error(w, fmt.Sprintf("Error al escribir la respuesta al cliente: %v", err), http.StatusInternalServerError)
            return
        }
    })

    // Iniciar el servidor HTTP
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    log.Printf("Servidor escuchando en el puerto %s...", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}

