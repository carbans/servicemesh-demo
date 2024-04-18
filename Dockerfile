# Usar una imagen base de Golang
FROM golang:latest

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar el código fuente al directorio de trabajo del contenedor
COPY . .

# Compilar la aplicación Go
RUN go build -o main .

# Exponer el puerto en el que el servidor HTTP escucha
EXPOSE 8080

# Comando para ejecutar la aplicación una vez que se inicie el contenedor
CMD ["./main"]

