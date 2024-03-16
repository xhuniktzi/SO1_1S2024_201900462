
import './App.css';
import TreeProcess from './TreeProcess';
import RamUsage from './RamUsage';
import CpuUsage from './CpuUsage';
import ChartComponent from './Historic';

function App() {
  return (
    <>
      <TreeProcess />
      <ChartComponent />

      <div className="wrapper">
        <div className="chartsjs">
          <h1>CPU</h1>
          <CpuUsage />
        </div>
        
        <div className="chartsjs">
          <h1>RAM</h1>
          <RamUsage />
        </div>
      </div>
    </>
  );
}

export default App;
