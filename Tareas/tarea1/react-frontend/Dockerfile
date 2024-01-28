# Usar la imagen oficial de Node.js para compilar el proyecto React
FROM node:16 as build

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar los archivos de definición de dependencias
COPY package.json ./
COPY package-lock.json ./

# Instalar dependencias del proyecto
RUN npm install

# Copiar el resto del código fuente del proyecto
COPY . ./

# Construir la aplicación React
RUN npm run build

# Usar Nginx para servir la aplicación
FROM nginx:stable-alpine

# Establecer la zona horaria
ENV TZ America/Guatemala

# Copiar el build de React al directorio de Nginx
COPY --from=build /app/build /usr/share/nginx/html

# Exponer el puerto 80
EXPOSE 80

# Comando por defecto para iniciar Nginx
CMD ["nginx", "-g", "daemon off;"]
