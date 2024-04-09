# Fase de construcción: Compila el código Rust en WebAssembly
FROM rust:1.73 as builder

# Crea el directorio de trabajo
WORKDIR /usr/src/app

# Copia los archivos de configuración de Rust y el código fuente al contenedor
COPY ./ .

# Construye el proyecto en modo release
RUN cargo build --release

# Fase de ejecución: Usa un entorno básico para ejecutar
FROM alpine:latest

# Instala dependencias necesarias en el entorno de ejecución
RUN apk --no-cache add ca-certificates

# Crea el directorio de trabajo
WORKDIR /root/

# Copia el binario compilado desde el builder al directorio de trabajo
COPY --from=builder /usr/src/app/target/release/wasm-producer .

# Expone el puerto necesario para el servidor HTTP
EXPOSE 8081

# Comando para ejecutar la aplicación
CMD ["./wasm-producer"]
