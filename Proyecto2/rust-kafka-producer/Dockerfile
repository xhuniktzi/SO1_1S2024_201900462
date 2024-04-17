# Usar una imagen de Rust oficial como etapa de compilación
FROM rust:1.67 as builder
WORKDIR /usr/src/app
COPY . .
RUN cargo build --release

# Etapa final
FROM debian:buster-slim
RUN apt-get update && apt-get install -y libssl1.1 ca-certificates && rm -rf /var/lib/apt/lists/*
COPY --from=builder /usr/src/app/target/release/rust_producer /usr/local/bin/rust_producer
CMD ["rust_producer"]
