// ImageUploader.js
import React, { useState } from 'react';

function ImageUploader() {
    const [selectedFile, setSelectedFile] = useState(null);

    const convertToBase64 = (file) => {
        return new Promise((resolve, reject) => {
            const reader = new FileReader();
            reader.readAsDataURL(file);
            reader.onload = () => resolve(reader.result);
            reader.onerror = (error) => reject(error);
        });
    };

    const onFileChange = (event) => {
        setSelectedFile(event.target.files[0]);
    };

    const onUpload = async () => {
        if (!selectedFile) {
            alert('Por favor, selecciona un archivo primero.');
            return;
        }

        const base64 = await convertToBase64(selectedFile);

        const url = 'http://localhost:3001/insertar'; // Sustituye esto con la URL real de tu API
        const body = {
            foto_b64: base64
        };

        try {
            const response = await fetch(url, {
                method: 'POST', // o 'PUT'
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(body),
            });

            if (response.ok) {
                const jsonResponse = await response.json();
                console.log('Respuesta de la API:', jsonResponse);
                alert('Imagen subida con éxito.');
            } else {
                console.error('Error en la respuesta de la API:', response.statusText);
                alert('Error al subir imagen.');
            }
        } catch (error) {
            console.error('Error al enviar la solicitud a la API:', error);
            alert('Error al subir imagen.');
        }
    };

    return (
        <div>
            <input type="file" onChange={onFileChange} accept="image/*" />
            <button onClick={onUpload}>Subir Imagen</button>
        </div>
    );
}

export default ImageUploader;
