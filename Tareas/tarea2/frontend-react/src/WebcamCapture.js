import React, { useRef } from 'react';
import Webcam from "react-webcam";

const WebcamCapture = () => {
    const webcamRef = useRef(null);

    const capture = React.useCallback(async () => {
        const imageSrc = webcamRef.current.getScreenshot();
        console.log(imageSrc); // Aquí tienes la imagen en base64
        // Puedes hacer algo con la imagen aquí, como enviarla a un servidor o mostrarla en la UI

        const url = 'http://localhost:3000/insertar'; // Sustituye esto con la URL real de tu API
        const body = {
            foto_b64: imageSrc
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

    }, [webcamRef]);

    return (
        <>
            <Webcam
                audio={false}
                ref={webcamRef}
                screenshotFormat="image/jpeg"
                width="100%"
            // Puedes especificar otras props aquí como la resolución de la cámara
            />
            <button onClick={capture}>Capturar Foto</button>
        </>
    );
};

export default WebcamCapture;
