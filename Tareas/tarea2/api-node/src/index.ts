import express, { Request, Response } from 'express';
import { MongoClient } from 'mongodb';
import cors from 'cors';
import bodyParser from 'body-parser';

const app = express();
const PORT = 3000;
const MONGO_URL = 'mongodb://mongo:27017';
const mongo_client = new MongoClient(MONGO_URL);
const DB_NAME = 'sopes1';

app.use(express.json({ limit: '50mb'}));
app.use(express.urlencoded({
    limit: '50mb',
    extended: true,
    parameterLimit: 50000
  }));

app.use(bodyParser.json({ limit: '50mb' }));
app.use(bodyParser.urlencoded({
    limit: '50mb',
    extended: true,
    parameterLimit: 50000
}));
app.use(cors());

app.get('/', (req: Request, res: Response) => {
    res.send('Hola mundo con TypeScript y Express!');
});

app.post('/insertar', async (req: Request, res: Response) => {
    try {
        await mongo_client.connect();
        const db = mongo_client.db(DB_NAME);
        const collection = db.collection('tarea2');

        const document = {
            foto_b64: req.body.foto_b64,
            fecha: new Date(),
        };

        const result = await collection.insertOne(document);
        res.json({ message: 'Foto insertada correctamente', result });
    } catch (error: unknown) {
        console.error(error);
        res.json({ message: 'Error al insertar la foto' });
    }
    finally {
        await mongo_client.close();
    }
});

app.get('/leer', async (req: Request, res: Response) => {
    try {
        await mongo_client.connect();
        const db = mongo_client.db(DB_NAME);
        const collection = db.collection('tarea2');

        const result = await collection.find().toArray();
        res.json(result);
    } catch (error: unknown) {
        res.json({ message: 'Error al leer las fotos' });
    }
    finally {
        await mongo_client.close();
    }
});

app.listen(PORT, () => {
    console.log(`Servidor corriendo en http://localhost:${PORT}`);
});
