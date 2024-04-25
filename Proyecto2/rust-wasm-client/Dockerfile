FROM rust:latest as builder
WORKDIR /usr/src/app
COPY . .
RUN cargo build --release
FROM ubuntu:latest
RUN apt-get update && apt-get install -y libssl-dev
COPY --from=builder /usr/src/app/target/release/rust_client /usr/local/bin/rust_client
EXPOSE 8000
CMD ["rust_client"]

# docker build -t desconocido502/rust_client .
# docker run -p 8000:8000 desconocido502/rust_client