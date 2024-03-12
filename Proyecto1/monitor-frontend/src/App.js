
import './App.css';
import TreeProcess from './TreeProcess';
import RamUsage from './RamUsage';
import CpuUsage from './CpuUsage';

function App() {
  return (
    <>
      <TreeProcess />
      <div class="wrapper">
        <div class="chartsjs">
          <h1>CPU</h1>
          <CpuUsage />
        </div>
        <div class="chartsjs">
          <h1>RAM</h1>
          <RamUsage />
        </div>
      </div>
    </>
  );
}

export default App;
