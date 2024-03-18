# Documentación Ampliada del Proyecto "Monitor de Recursos del Sistema" con Nuevas Especificaciones Técnicas

## Descripción General

El proyecto "Monitor de Recursos del Sistema" es una aplicación web desarrollada para monitorear y mostrar en tiempo real el uso de los recursos del sistema, enfocándose en la CPU y la RAM. Utiliza una arquitectura de microservicios contenerizados para garantizar una implementación, escalabilidad y mantenimiento eficientes.

## Arquitectura y Componentes

### Docker y Docker Compose

- **Docker:** Facilita la creación de contenedores para cada servicio del proyecto, asegurando un entorno de ejecución consistente.
- **Docker Compose:** Define y orquesta los servicios del proyecto, incluyendo la API en Go, el frontend en React y la base de datos MySQL, a través de un archivo YAML.

```yaml
version: '3.8'

services:
  api:
    build: ./broker-backend
    ports:
      - "8080:8080"
    networks:
      - mynetwork
    depends_on:
      - mysql

  frontend:
    build: ./monitor-frontend
    ports:
      - "80:80"
    depends_on:
      - api
    networks:
      - mynetwork

  mysql:
    image: mysql
    command: --default-authentication-plugin=mysql_native_password
    volumes:
      - mysql-data:/var/lib/mysql
    ports:
      - "3306:3306"
    environment:
      MYSQL_ROOT_PASSWORD: rootpassword
      MYSQL_DATABASE: monitor
      MYSQL_USER: user
      MYSQL_PASSWORD: password
    networks:
      - mynetwork

volumes:
  mysql-data:

networks:
  mynetwork:
    external: true
```

### Servicios Definidos en Docker Compose

1. **API (Backend en Go):** Realiza la lógica de negocios, interactúa con la base de datos MySQL y los módulos del kernel.
2. **Frontend (React):** Muestra la interfaz de usuario con gráficos en tiempo real, utilizando datos de la API.
3. **MySQL:** Almacena información histórica sobre el uso de recursos, permitiendo análisis y visualización de tendencias.

### Construcción y Despliegue de la API

- **Imagen Base:** Se utiliza `golang:1.18` como imagen base para compilar la aplicación Go.
- **Compilación:** Se copian los archivos `go.mod` y `go.sum`, se descargan las dependencias y se compila la aplicación.
- **Imagen Final:** Se usa `alpine:latest` como imagen base para la imagen final, donde se copia el ejecutable compilado.

### Backend (API en Go)

- **Inicialización:** Se abre la conexión a la base de datos y se inicia un ticker que inserta datos cada 10 segundos.
- **Controladores:** Existen controladores específicos para obtener datos de la CPU y la RAM, y para recuperar los últimos registros de las bases de datos.

```go
package main

import (
 "broker-backend/config"
 "broker-backend/routes"
 "log"
 "net/http"
)

func main() {
 config.OpenDB()
 go insertarDatosPeriodicamente() // Función para insertar datos

 router := mux.NewRouter()
 routes.InitializeRoutes(router)

 log.Println("Servidor corriendo en el puerto 8080")
 http.ListenAndServe(":8080", router)
}

func insertarDatosPeriodicamente() {
 // Implementación para insertar datos periódicamente
}
```

### Frontend (React)

- **Construcción:** Utiliza `node:16` como imagen base para compilar la aplicación React.
- **Nginx:** Se emplea `nginx:stable-alpine` para servir la aplicación React, configurando Nginx para redirigir solicitudes a la API.

```javascript
import { useEffect, useState } from 'react';

const CpuUsage = () => {
    const [cpuUsage, setCpuUsage] = useState({ used: 0, free: 0 });

    useEffect(() => {
        fetch('/api/cpu')
            .then(response => response.json())
            .then(data => setCpuUsage({ used: data.used, free: data.free }));
    }, []);

    return (
        <div>
            <p>CPU Used: {cpuUsage.used}%</p>
            <p>CPU Free: {cpuUsage.free}%</p>
        </div>
    );
};

export default CpuUsage;
```

### Nginx

- **Configuración:** Se define una configuración específica para redirigir las solicitudes a los endpoints de la API y servir la aplicación React.

```nginx
server {
    listen 80;
    server_name localhost;

    location / {
        root /usr/share/nginx/html;
        index index.html;
    }

    location /api {
        proxy_pass http://api:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

### Módulos del Kernel de Linux

- **Implementación:** Se detalla la implementación de los módulos del kernel para la CPU y la RAM, mostrando cómo se recopilan y exponen los datos.

### Base de Datos (MySQL)

- **Esquemas:** Se crean tablas `cpu` y `ram` para almacenar los registros de uso de recursos.
- **Inicialización:** Se configura el usuario y la contraseña de MySQL, y se inicializa la base de datos `monitor`.

### Comunicación entre Componentes

- **Frontend - Backend:** Utiliza solicitudes HTTP para comunicarse.
- **Backend - Módulos del Kernel:** El backend lee archivos en `/proc` para obtener los datos.
- **Backend - Base de Datos:** Inserta y recupera datos de uso de recursos.

## Instrucciones de Uso

1. **Inicialización:** Ejecutar `docker-compose up` para levantar los servicios.
2. **Acceso al Frontend:** Acceder a `http://localhost` para visualizar la interfaz de usuario.
3. **API:** Interactuar con endpoints como `/cpu`, `/ram`, y `/data` para obtener datos en tiempo real y registros históricos.
