# SubServer.py
import redis

# Asumiendo que tienes las credenciales de tu instancia de MemoryStore de Redis
REDIS_PORT = 6379  # Puerto por defecto de Redis
REDIS_HOST = '127.0.0.1'  # Host de la instancia de Redis

# Conectar al servidor Redis
sub_client = redis.StrictRedis(host=REDIS_HOST, port=REDIS_PORT, decode_responses=True)

# Crear un suscriptor
pubsub = sub_client.pubsub()

# Suscribirse al canal 'test'
pubsub.subscribe('test')

print('Subscribed to "test" channel. Waiting for messages...')

# Escuchar mensajes
for message in pubsub.listen():
    if message['type'] == 'message':
        print('Received message:', message['data'])