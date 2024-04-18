# Etapa de construcción
FROM node:16 as build-stage

# Directorio de trabajo
WORKDIR /app

# Instalar dependencias
COPY package*.json ./
RUN npm install

# Construir aplicación
COPY . .
RUN npm run build

# Etapa de producción con Nginx
FROM nginx:stable-alpine as production-stage

# Copiar la configuración personalizada de Nginx
COPY nginx.conf /etc/nginx/nginx.conf

# Copiar archivos estáticos construidos
COPY --from=build-stage /app/dist /usr/share/nginx/html

# Exponer el puerto 80 para el acceso HTTP
EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
