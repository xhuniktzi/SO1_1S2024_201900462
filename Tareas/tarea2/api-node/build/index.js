"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const mongodb_1 = require("mongodb");
const cors_1 = __importDefault(require("cors"));
const body_parser_1 = __importDefault(require("body-parser"));
const app = (0, express_1.default)();
const PORT = 3000;
const MONGO_URL = 'mongodb://localhost:27017';
const mongo_client = new mongodb_1.MongoClient(MONGO_URL);
const DB_NAME = 'sopes1';
app.use(express_1.default.json({ limit: '50mb' }));
app.use(express_1.default.urlencoded({
    limit: '50mb',
    extended: true,
    parameterLimit: 50000
}));
app.use(body_parser_1.default.json({ limit: '50mb' }));
app.use(body_parser_1.default.urlencoded({
    limit: '50mb',
    extended: true,
    parameterLimit: 50000
}));
app.use((0, cors_1.default)());
app.get('/', (req, res) => {
    res.send('Hola mundo con TypeScript y Express!');
});
app.post('/insertar', (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    try {
        yield mongo_client.connect();
        const db = mongo_client.db(DB_NAME);
        const collection = db.collection('tarea2');
        const document = {
            foto_b64: req.body.foto_b64,
            fecha: new Date(),
        };
        const result = yield collection.insertOne(document);
        res.json({ message: 'Foto insertada correctamente', result });
    }
    catch (error) {
        console.error(error);
        res.json({ message: 'Error al insertar la foto' });
    }
    finally {
        yield mongo_client.close();
    }
}));
app.get('/leer', (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    try {
        yield mongo_client.connect();
        const db = mongo_client.db(DB_NAME);
        const collection = db.collection('tarea2');
        const result = yield collection.find().toArray();
        res.json(result);
    }
    catch (error) {
        res.json({ message: 'Error al leer las fotos' });
    }
    finally {
        yield mongo_client.close();
    }
}));
app.listen(PORT, () => {
    console.log(`Servidor corriendo en http://localhost:${PORT}`);
});
