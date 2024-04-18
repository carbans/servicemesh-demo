package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		targetEndpoint := os.Getenv("TARGET_ENDPOINT")
		if targetEndpoint == "" {
			http.Error(w, "TARGET_ENDPOINT no está configurado", http.StatusInternalServerError)
			return
		}

		resp, err := http.Get(targetEndpoint)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error al realizar la solicitud al endpoint de destino: %v", err), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			http.Error(w, fmt.Sprintf("El endpoint de destino respondió con un código de estado no válido: %d", resp.StatusCode), http.StatusInternalServerError)
			return
		}

		// Copiar encabezados de la respuesta del endpoint de destino a la respuesta al cliente
		for key, value := range resp.Header {
			w.Header()[key] = value
		}

		// Copiar el cuerpo de la respuesta del endpoint de destino a la respuesta al cliente
		w.WriteHeader(resp.StatusCode)
		_, err = io.Copy(w, resp.Body)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error al escribir la respuesta al cliente: %v", err), http.StatusInternalServerError)
			return
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Servidor escuchando en el puerto %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
