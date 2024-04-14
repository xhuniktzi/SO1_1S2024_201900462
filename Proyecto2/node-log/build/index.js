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
const dotenv_1 = __importDefault(require("dotenv"));
const cors_1 = __importDefault(require("cors"));
// Configura dotenv
dotenv_1.default.config();
const app = (0, express_1.default)();
const port = process.env.PORT || 3000; // Permite configurar el puerto también
const uri = process.env.MONGO_URI || "mongodb://localhost:27017"; // Utiliza valor predeterminado si no se define en .env
app.use((0, cors_1.default)());
// Cliente de MongoDB
const client = new mongodb_1.MongoClient(uri);
app.get('/logs', (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    try {
        yield client.connect();
        const database = client.db("logs");
        const logs = database.collection("vote-logs");
        // Recupera los últimos 20 logs
        const query = logs.find({}).sort({ timestamp: -1 }).limit(20);
        const results = yield query.toArray();
        res.status(200).json(results);
    }
    catch (e) {
        const error = e;
        res.status(500).json({ message: error.message });
    }
    finally {
        yield client.close();
    }
}));
app.listen(port, () => {
    console.log(`Server running at http://localhost:${port}`);
});
