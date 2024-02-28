# Usar una imagen base de Go
FROM golang:1.18-alpine as builder

# Establecer el directorio de trabajo
WORKDIR /app

# Copia el código fuente del proyecto al contenedor
COPY . .

# Ejecutar el script que inicializa el módulo de Go si es necesario
RUN chmod +x ./init-module.sh && ./init-module.sh

# Continuar con los pasos de construcción, como compilar tu aplicación Go
RUN go build -o main .

# Usar una imagen limpia para el runtime
FROM alpine:latest  
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
