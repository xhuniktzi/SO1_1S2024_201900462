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

# Instala Nginx y gettext (para envsubst)
# RUN apk add --no-cache nginx gettext

# Crea el directorio para los archivos de PID y de sesión de Nginx
# RUN mkdir -p /run/nginx

# Copiar la configuración personalizada de Nginx
# COPY nginx.conf /etc/nginx/nginx.conf.template

# Copiar archivos estáticos construidos
COPY --from=build-stage /app/dist /usr/share/nginx/html

# Crea un script para iniciar Nginx después de aplicar envsubst
# RUN echo "envsubst '\${PROXY_URL}' < /etc/nginx/nginx.conf.template > /etc/nginx/nginx.conf && nginx -g 'daemon off;'" > /start.sh

# RUN chmod +x /start.sh

# Exponer el puerto 80 para el acceso HTTP
EXPOSE 80

# Iniciar Nginx
CMD ["nginx", "-g", "daemon off;"]
