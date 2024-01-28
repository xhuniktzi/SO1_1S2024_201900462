import React, { useState } from 'react';
import './App.css';

function App() {
  const [studentData, setStudentData] = useState(null);

  const fetchData = async () => {
    const response = await fetch('http://localhost:8080/data');
    const data = await response.json();
    setStudentData(data);
  };

  return (
    <div className="App">
      <header className="App-header">
        <p>
          Pulsa el botón para obtener datos del estudiante.
        </p>
        <button onClick={fetchData}>
          Obtener Datos
        </button>
        {studentData && <div>
          <p>Carnet: {studentData.carnet}</p>
          <p>Nombre: {studentData.nombre}</p>
          <p>Hora: {new Date().toISOString()}</p>
        </div>}
      </header>
    </div>
  );
}

export default App;
