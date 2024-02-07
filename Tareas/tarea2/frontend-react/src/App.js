
import './App.css';
import ImageDisplay from './ImageDisplay';
import WebcamCapture from './WebcamCapture';


function App() {
  return (
    <div className="App">
      <h1>Subir y mostrar imagen con React</h1>
      {/* <ImageUploader /> */}
      <WebcamCapture />
      <ImageDisplay imageUrl="http://localhost:3000/leer" />
    </div>
  );
}

export default App;
