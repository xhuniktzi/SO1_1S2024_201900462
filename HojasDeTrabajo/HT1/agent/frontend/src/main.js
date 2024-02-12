import './style.css';
import './app.css';

import {ObtenerMemoria} from '../wailsjs/go/main/App';

const obtenerMemoria = async () => {
    const response = await ObtenerMemoria();
    
    const parsed = JSON.parse(response);
    
    document.querySelector('#free_mem').innerHTML = parsed.Libre;
    document.querySelector('#total_mem').innerHTML = parsed.Total;
}

document.querySelector('#app').innerHTML = `
    <p id="free_mem"></p>
    <p id="total_mem"></p>
`

setInterval(() => {
    obtenerMemoria();
}, 1000);