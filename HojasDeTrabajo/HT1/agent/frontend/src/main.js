import './style.css';
import './app.css';

// import logo from './assets/images/logo-universal.png';
import {Greet, ObtenerMemoria} from '../wailsjs/go/main/App';

const obtenerMemoria = async () => {
    const response = await ObtenerMemoria();
    console.log(response);
    document.querySelector('#response_mem').innerHTML = response;    
}

document.querySelector('#app').innerHTML = `
    <p id="response_mem"></p>
`

setInterval(() => {
    obtenerMemoria();
}, 5000);

// document.querySelector('#app').innerHTML = `
//     <img id="logo" class="logo">
//       <div class="result" id="result">Please enter your name below 👇</div>
//       <div class="input-box" id="input">
//         <input class="input" id="name" type="text" autocomplete="off" />
//         <button class="btn" onclick="greet()">Greet</button>
//       </div>
//     </div>
// `;
// document.getElementById('logo').src = logo;

// let nameElement = document.getElementById("name");
// nameElement.focus();
// let resultElement = document.getElementById("result");

// Setup the greet function
// window.greet = function () {
//     // Get name
//     let name = nameElement.value;

//     // Check if the input is empty
//     if (name === "") return;

//     // Call App.Greet(name)
//     try {
//         Greet(name)
//             .then((result) => {
//                 // Update result with data back from App.Greet()
//                 resultElement.innerText = result;
//             })
//             .catch((err) => {
//                 console.error(err);
//             });
//     } catch (err) {
//         console.error(err);
//     }
// };

