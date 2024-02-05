
import './App.css';
import ImageDisplay from './ImageDisplay';
import ImageUploader from './ImageUploader';


function App() {
  return (
    <div className="App">
      <h1>Subir y mostrar imagen con React</h1>
      <ImageUploader />
      <ImageDisplay imageUrl="http://localhost:3001/leer" />
    </div>
    );
}

export default App;
