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

# Establecer la zona horaria
ENV TZ America/Guatemala

# Exponer el puerto 3000
EXPOSE 3000

# Comando para ejecutar la aplicación
CMD ["node", "build/index.js"]