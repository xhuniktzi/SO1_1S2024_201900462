import { React, useEffect, useState, useRef } from 'react'
import { Chart, PieController, ArcElement, Tooltip, Legend } from 'chart.js';
Chart.register(PieController, ArcElement, Tooltip, Legend);

const RamUsage = () => {
    const [ramUsage, setRamUsage] = useState({ used: 0, free: 0 });
    const chartRef = useRef(null); // Referencia para almacenar la instancia del gráfico

    // Obtener datos de RAM cada 500 ms
    useEffect(() => {
        const interval = setInterval(() => {
            fetch('/ram')
                .then(response => response.json())
                .then(data => {
                    setRamUsage({ used: data.UsagePercent, free: 100 - data.UsagePercent });
                });
        }, 1000);

        // Limpieza del intervalo
        return () => clearInterval(interval);
    }, []);

    useEffect(() => {
        const ctx = document.getElementById('ram-usage').getContext('2d');
        const data = {
            labels: ['Used', 'Free'],
            datasets: [{
                label: 'RAM Usage',
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
    }, []); // Dependencia en ramUsage

    useEffect(() => {
        if (chartRef.current) {
            chartRef.current.data.datasets.forEach((dataset) => {
                dataset.data = [ramUsage.used, ramUsage.free];
            });
            chartRef.current.update();
        }
    }, [ramUsage]);

    return (
        <canvas id="ram-usage" width="200" height="200"></canvas>
    )
}

export default RamUsage;
