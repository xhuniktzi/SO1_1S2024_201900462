import React, { useEffect, useState } from 'react';

function ImageDisplay({ imageUrl }) {
    const [imagesSrc, setImagesSrc] = useState([]);

    useEffect(() => {
        fetch(imageUrl)
            .then(response => response.json())
            .then(data => {
                // Mapear el arreglo de objetos a un arreglo de cadenas Base64
                const imagesBase64 = data.map(item => item.foto_b64);
                setImagesSrc(imagesBase64);
            });
    }, [imageUrl]);

    return (
        <div>
            {imagesSrc.map((src, index) => (
                <img key={index} src={`${src}`} alt={`Loaded from API ${index}`} />
            ))}
        </div>
    );
}

export default ImageDisplay;
