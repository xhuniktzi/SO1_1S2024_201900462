# Usar una imagen base oficial de Node
FROM node:16-slim as build-stage

# Crear directorio de trabajo
WORKDIR /usr/src/app

# Copiar 'package.json' y 'package-lock.json' (si están disponibles)
COPY package*.json ./

# Instalar dependencias incluyendo las de desarrollo para la compilación
RUN npm install

# Copiar los archivos de proyecto
COPY . .

# Compilar TypeScript a JavaScript
RUN npm run build

# Etapa de producción, usar una nueva etapa para mantener la imagen ligera
FROM node:16-slim as production-stage

# Crear directorio de trabajo
WORKDIR /usr/src/app

# Instalar solo las dependencias de producción
COPY package*.json ./
RUN npm install --only=production

# Copiar los archivos compilados de la etapa de construcción
COPY --from=build-stage /usr/src/app/build ./build

# Exponer el puerto que tu aplicación usará
EXPOSE 3000

# Comando para correr la aplicación
CMD ["node", "build/index.js"]
