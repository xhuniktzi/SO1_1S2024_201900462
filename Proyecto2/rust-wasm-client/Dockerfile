# Etapa de compilación para WebAssembly
FROM rust:1.73 as builder
WORKDIR /usr/src/app

# Añadir el target para WebAssembly
RUN rustup target add wasm32-unknown-unknown

# Instalar wasm-bindgen CLI
# Ensure this version matches the one in Cargo.toml
ARG WASM_BINDGEN_VERSION=0.2.78
RUN cargo install wasm-bindgen-cli --vers $WASM_BINDGEN_VERSION

# Copiar todos los archivos necesarios para la compilación
COPY . .

# Compilar el proyecto para WebAssembly
RUN cargo build --target wasm32-unknown-unknown --release

# Usar wasm-bindgen para generar los archivos necesarios para el navegador
RUN wasm-bindgen target/wasm32-unknown-unknown/release/wasm_client.wasm --out-dir ./pkg --web

# Etapa para servir los archivos con nginx
FROM nginx:alpine
COPY --from=builder /usr/src/app/pkg /usr/share/nginx/html
COPY my_nginx.conf /etc/nginx/nginx.conf
