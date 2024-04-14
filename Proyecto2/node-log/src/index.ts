import express from 'express';
import { MongoClient } from 'mongodb';
import dotenv from 'dotenv';

// Configura dotenv
dotenv.config();

const app = express();
const port = process.env.PORT || 3000; // Permite configurar el puerto también
const uri = process.env.MONGO_URI || "mongodb://localhost:27017"; // Utiliza valor predeterminado si no se define en .env

// Cliente de MongoDB
const client = new MongoClient(uri);

app.get('/logs', async (req, res) => {
    try {
        await client.connect();
        const database = client.db("logs");
        const logs = database.collection("vote-logs");

        // Recupera los últimos 20 logs
        const query = logs.find({}).sort({ timestamp: -1 }).limit(20);
        const results = await query.toArray();

        res.status(200).json(results);
    } catch (e) {
        const error = e as Error;
        res.status(500).json({ message: error.message });
    } finally {
        await client.close();
    }
});

app.listen(port, () => {
    console.log(`Server running at http://localhost:${port}`);
});
