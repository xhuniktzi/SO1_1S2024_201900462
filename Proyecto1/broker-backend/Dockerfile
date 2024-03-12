# Usar la imagen oficial de Go como imagen base
FROM golang:1.18 as builder

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar los archivos Go mod y sum
COPY go.mod ./
COPY go.sum ./

# Descargar las dependencias de Go
RUN go mod download

# Copiar el código fuente del proyecto Go al directorio de trabajo
COPY . .

# Compilar la aplicación Go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Usar una imagen base mínima para la imagen final
FROM alpine:latest  

# Establecer la zona horaria
ENV TZ America/Guatemala

# Establecer el directorio de trabajo
WORKDIR /root/

# Copiar el ejecutable compilado desde la etapa de compilación
COPY --from=builder /app/main .

# Exponer el puerto en el que se ejecutará el servidor
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./main"]