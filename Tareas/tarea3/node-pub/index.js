// PubServer.js
const redis = require('redis');

// Asumiendo que tienes las credenciales de tu instancia de MemoryStore de Redis
const REDIS_PORT = 6379; // Puerto por defecto de Redis
const REDIS_HOST = '127.0.0.1'; // Host de la instancia de Redis


// Crear un cliente Redis
const pubClient = redis.createClient({
url: 'redis://' + REDIS_HOST + ':' + REDIS_PORT
});

pubClient.on('error', (err) => console.log('Redis Client Error', err));

pubClient.on('connect', () => {
  console.log('Connected to Redis Server!');

  // Publicar mensaje en el canal 'test'
  setInterval(() => {
    const message = JSON.stringify({ msg: "Hola a todos" });

    pubClient.publish('test', message, (err, reply) => {
      if (err) {
        console.log('Error publishing message:', err);
      } else {
        console.log(`Message published to channel 'test': ${message}`);
      }
    });
  }, 5000); // Publicar cada 5 segundos
});

// Agregar manejo de error
pubClient.on('error', (err) => {
  console.error('Redis Client Error', err);
  process.exit(1); // Salir del proceso en caso de error para no continuar
});

pubClient.connect();