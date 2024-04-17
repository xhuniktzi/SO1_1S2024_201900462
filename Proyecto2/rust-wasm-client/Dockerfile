# Etapa de compilación para WebAssembly
FROM rust:1.62 as builder
WORKDIR /usr/src/app
COPY . .
RUN cargo build --target wasm32-unknown-unknown --release
RUN wasm-bindgen target/wasm32-unknown-unknown/release/wasm_client.wasm --out-dir ./pkg --web

# Servir con nginx
FROM nginx:alpine
COPY --from=builder /usr/src/app/pkg /usr/share/nginx/html
COPY my_nginx.conf /etc/nginx/nginx.conf
