import './style.css';
import './app.css';

import {ObtenerMemoria} from '../wailsjs/go/main/App';

document.querySelector('#app').innerHTML = `
<canvas id="miGraficoPastel"></canvas>
`

let miGrafico = document.getElementById('miGraficoPastel').getContext('2d');
let graficoPastel = new Chart(miGrafico, {
    type: 'pie', // Tipo de gráfico
    data: {
        labels: ['LIBRE', 'EN USO'], // Etiquetas de los datos
        datasets: [{
            label: 'Uso de memoria',
            data: [10, 20], // Datos iniciales
            backgroundColor: [
                'rgba(255, 99, 132, 0.2)',
                'rgba(54, 162, 235, 0.2)',
                'rgba(255, 206, 86, 0.2)'
            ],
            borderColor: [
                'rgba(255, 99, 132, 1)',
                'rgba(54, 162, 235, 1)',
                'rgba(255, 206, 86, 1)'
            ],
            borderWidth: 1
        }]
    },
    options: {}
});

const obtenerMemoria = async () => {
    const response = await ObtenerMemoria();
    
    let parsed = JSON.parse(response);
    parsed.Total = parsed.Total - parsed.Libre;
    const values = Object.values(parsed);

    graficoPastel.data.datasets.forEach(dataset => {
        dataset.data = values;
    });

    graficoPastel.update();
}

setInterval(() => {
    obtenerMemoria();
}, 500);