package main

import (
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

		// Crear una nueva solicitud HTTP para el endpoint de destino
		req, err := http.NewRequest(r.Method, targetEndpoint, r.Body)
		if err != nil {
			http.Error(w, "Error creando la solicitud", http.StatusInternalServerError)
			return
		}

		// Copiar los encabezados originales de la solicitud
		for header, values := range r.Header {
			for _, value := range values {
				req.Header.Add(header, value)
			}
		}

		// Hacer la solicitud al endpoint de destino
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, "Error en la solicitud al endpoint de destino", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Copiar los encabezados de la respuesta del endpoint de destino
		for header, values := range resp.Header {
			for _, value := range values {
				w.Header().Add(header, value)
			}
		}

		// Establecer el código de estado de la respuesta
		w.WriteHeader(resp.StatusCode)

		// Copiar el cuerpo de la respuesta del endpoint de destino
		_, err = io.Copy(w, resp.Body)
		if err != nil {
			http.Error(w, "Error al copiar el cuerpo de la respuesta", http.StatusInternalServerError)
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
