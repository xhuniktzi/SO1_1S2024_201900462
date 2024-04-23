
# 1. Introducción

Este documento describe la arquitectura y el funcionamiento de una aplicación distribuida basada en microservicios que interactúan a través de gRPC, Kafka y Redis, desplegada en Kubernetes. La aplicación está diseñada para manejar votaciones de música, almacenando y procesando información relacionada con álbumes, artistas y votaciones.

# 2. Objetivos

- **Desarrollar un sistema robusto** para el procesamiento de votaciones en tiempo real.
- **Utilizar tecnologías de contenedores** para facilitar el despliegue y escalabilidad del sistema.
- **Implementar microservicios** que interactúen eficientemente utilizando gRPC y Kafka.
- **Garantizar la persistencia de datos** y su rápida recuperación mediante bases de datos como MongoDB y Redis.

# 3. Descripción de cada tecnología utilizada

- **Kafka**: Utilizado para manejar el flujo de mensajes entre los servicios. Kafka permite la transmisión de datos en tiempo real entre productores y consumidores de manera eficiente y escalable.
- **Redis**: Base de datos en memoria que ofrece almacenamiento rápido de clave-valor, utilizada para almacenar votaciones temporalmente antes de ser procesadas.
- **MongoDB**: Base de datos NoSQL que se utiliza para almacenar logs de operaciones y errores, aprovechando su esquema flexible y capacidad de manejo de grandes volúmenes de datos.
- **gRPC**: Protocolo de comunicación entre servicios, usado para las interacciones entre el cliente y el servidor. Proporciona una comunicación rápida y eficiente.
- **Docker**: Tecnología de contenerización que facilita el despliegue y la gestión de microservicios, asegurando la portabilidad y consistencia del entorno de ejecución.
- **Node.js**: Plataforma de ejecución para JavaScript en el servidor, utilizada aquí para desarrollar servicios backend que se comunican con MongoDB y proporcionan una API para el frontend.
- **Vue.js**: Framework de JavaScript para construir interfaces de usuario. En este caso, se utiliza para desarrollar el frontend que interactúa con el servicio de Node.js para mostrar los logs de MongoDB.
- **Locust**: Herramienta de prueba de carga de código abierto escrita en Python, utilizada para realizar pruebas de carga en el sistema y asegurar que pueda manejar el tráfico esperado en producción.

# 4. Descripción de cada deployment y service de K8S

## Deployments

- **MongoDB**: Base de datos principal para el almacenamiento de logs. Desplegado con un PersistentVolume para garantizar la durabilidad de los datos.
- **Redis**: Actúa como almacenamiento temporal para las votaciones antes de ser procesadas.
- **Grafana**: Herramienta de visualización y análisis de datos que consume datos de Prometheus.
- **Zookeeper**: Servicio de coordinación para Kafka que gestiona la configuración de la red y sincronización.
- **Kafka**: Sistema de colas de mensajes que facilita la transferencia de datos entre diferentes componentes del sistema.
- **Consumer**: Microservicio que consume mensajes de Kafka, procesa los datos y los almacena en Redis.
- **Golang-producer**: Microservicio que recibe peticiones y las envía a Kafka.
- **Golang-client**: Interfaz de usuario para enviar votaciones que luego son procesadas por otros servicios.

## Services

- Cada deployment está asociado con un service en Kubernetes que permite la comunicación entre los diferentes componentes, facilitando la exposición de los servicios a través de puertos definidos y realizando el balanceo de carga necesario.

## Otros Despliegues (fuera de Kubernetes)

- **Node-Log**: Servidor de backend desarrollado en Node.js que se ejecuta en Cloud Run. Este servicio maneja las solicitudes de log y se comunica con MongoDB para recuperar la información de logs, la cual es devuelta a la interfaz de usuario.
- **Vue-Log**: Aplicación de frontend desarrollada con Vue.js, también hospedada en Cloud Run. Esta interfaz de usuario consume datos de la API proporcionada por el servicio Node-Log, permitiendo a los usuarios visualizar y actualizar los registros de logs de forma dinámica.

## Ingress

- **Ingress para golang-client**: Proporciona reglas que permiten el acceso externo a los servicios dentro del cluster. Este Ingress en particular está configurado para redirigir tráfico a `golang-client` a través del path `/golang-client`, utilizando nginx como controlador de Ingress. Esto facilita el acceso público al servicio de interfaz de usuario que interactúa con el microservicio de gRPC, simplificando la gestión de tráfico y seguridad al nivel de la aplicación.

# Horizontal Pod Autoscaler

- **HPA para Consumer**: Automatiza el escalado de los pods del deployment `consumer` en función del uso de CPU. Está configurado para mantener la utilización de CPU cerca del 50%, escalando entre 2 y 5 replicas según sea necesario. Esto asegura que el servicio pueda manejar aumentos en la carga de trabajo sin degradar el rendimiento, optimizando el uso de recursos y mejorando la disponibilidad del servicio.

# 5. Ejemplo de funcionamiento con capturas de pantalla

- La interfaz de usuario realizando una votación.
- La visualización de los logs en MongoDB a través de la aplicación Node.js.
- Las métricas de Grafana mostrando el rendimiento de la aplicación.
- La interfaz de Vue.js interactuando con el servicio Node-Log para recuperar y mostrar logs almacenados en MongoDB.
- El panel de Cloud Run mostrando el estado de los servicios Node-Log y Vue-Log, indicando cómo se maneja la escalabilidad y el balanceo de carga de manera automática.

# 6. Conclusiones

- El uso de Kubernetes y Docker ha simplificado el despliegue y escalabilidad de la aplicación, permitiendo una gestión eficiente de los microservicios.
- Kafka y Redis se han integrado de manera efectiva para manejar grandes volúmenes de datos en tiempo real, asegurando un procesamiento rápido y fiable de las votaciones.
- La aplicación ha demostrado ser robusta y capaz de manejar picos de carga mediante el uso de la escalabilidad horizontal y el balanceo de carga proporcionado por Kubernetes.
- **Flexibilidad en la implementación**: El uso de Cloud Run para los servicios de Node.js y Vue.js demuestra una arquitectura flexible, aprovechando las capacidades de autoescalado y manejo de alto tráfico sin necesidad de gestionar la infraestructura subyacente.
- **Separación eficiente de responsabilidades**: Al desplegar los servicios front y back-end en Cloud Run, se facilita la gestión y actualización de cada componente de manera independiente, lo que puede resultar en ciclos de desarrollo y mantenimiento más rápidos y menos propensos a errores.

- **Gestión dinámica de la carga**: La implementación del Horizontal Pod Autoscaler permite que la aplicación se adapte dinámicamente a las variaciones en la carga de trabajo, manteniendo un rendimiento óptimo sin intervención manual. Esto es crucial para mantener la eficiencia operativa y minimizar los costos al ajustar los recursos de forma automática según la demanda.
- **Configuración de Ingress**: El uso de Ingress para exponer los servicios permite una configuración centralizada de las reglas de enrutamiento, lo que facilita la implementación de políticas de seguridad, balanceo de carga y cifrado SSL en un solo punto, simplificando la gestión de la red y mejorando la seguridad del sistema.

# 7. Referencias

- [Kubernetes](https://kubernetes.io/)
- [Docker](https://www.docker.com/)
- [gRPC](https://grpc.io/)
- [Kafka](https://kafka.apache.org/)
- [Redis](https://redis.io/)
- [MongoDB](https://www.mongodb.com/)
- [Node.js](https://nodejs.org/)
- [Vue.js](https://vuejs.org/)
- [Locust](https://locust.io/)
- [Cloud Run](https://cloud.google.com/run)
- [Grafana](https://grafana.com/)
