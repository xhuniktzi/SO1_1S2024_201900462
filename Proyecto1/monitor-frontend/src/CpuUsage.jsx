import { React, useEffect, useState, useRef} from 'react'
import { Chart, PieController, ArcElement, Tooltip, Legend } from 'chart.js';
Chart.register(PieController, ArcElement, Tooltip, Legend);

const CpuUsage = () => {
    const [cpuUsage, setCpuUsage] = useState({ used: 0, free: 0 });
    const chartRef = useRef(null); // Referencia para almacenar la instancia del gráfico

    // Obtener datos de CPU cada 500 ms
    useEffect(() => {
        const interval = setInterval(() => {
            fetch('http://localhost:8080/cpu')
                .then(response => response.json())
                .then(data => {
                    const cpu_cores = 2;
                    let usage = data.Total_CPU_Time / cpu_cores;;
                    if (usage > 1) {
                        usage = 100;
                    } else {
                        usage = usage * 100;
                    }
                    setCpuUsage({ used: usage, free: 100 - usage });
                });
        }, 1000);

        // Limpieza del intervalo
        return () => clearInterval(interval);
    }, []);

    useEffect(() => {
        const ctx = document.getElementById('cpu-usage').getContext('2d');
        const data = {
            labels: ['Used', 'Free'],
            datasets: [{
                label: 'CPU Usage',
                data: [50, 50], // Usar los valores del estado
                backgroundColor: [
                    'rgba(255, 99, 132, 0.2)',
                    'rgba(54, 162, 235, 0.2)'
                ],
                borderColor: [
                    'rgba(255, 99, 132, 1)',
                    'rgba(54, 162, 235, 1)'
                ],
                borderWidth: 1,
                hoverOffset: 4
            }]
        };

        if (!chartRef.current) {
            // Crear el nuevo gráfico y almacenar su referencia
            chartRef.current = new Chart(ctx, {
                type: 'pie',
                data: data,
            });
        }
    }, []); // Dependencia en cpuUsage


    useEffect(() => {
        if (chartRef.current) {
            chartRef.current.data.datasets.forEach((dataset) => {
                dataset.data = [cpuUsage.used, cpuUsage.free];
            });
            chartRef.current.update();
        }
    }, [cpuUsage]);

    return (
        <canvas id="cpu-usage" width="200" height="200"></canvas>
    )
};

export default CpuUsage;