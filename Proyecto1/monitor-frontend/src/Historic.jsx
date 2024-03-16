import './App.css';
import React, { useState, useEffect } from 'react';
import { Chart as ChartJS, CategoryScale, LinearScale, BarElement, Title, Tooltip, Legend } from 'chart.js';
import { Bar } from 'react-chartjs-2';

ChartJS.register(CategoryScale, LinearScale, BarElement, Title, Tooltip, Legend);

const ChartComponent = () => {
    const [data, setData] = useState({ cpu: [], ram: [] });

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await fetch('http://localhost:8080/data');
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                const json = await response.json();
                setData(json);
            } catch (error) {
                console.error("Could not fetch the data", error);
            }
        };

        fetchData();
        const interval = setInterval(fetchData, 500);

        return () => clearInterval(interval);
    }, []);

    const cpuData = {
        labels: data.cpu.map(item => `ID ${item.id}`),
        datasets: [
            {
                label: 'CPU Used (%)',
                data: data.cpu.map(item => item.used),
                backgroundColor: 'rgba(255, 99, 132, 0.5)',
            },
            {
                label: 'CPU Free (%)',
                data: data.cpu.map(item => item.free),
                backgroundColor: 'rgba(53, 162, 235, 0.5)',
            }
        ],
    };

    const ramData = {
        labels: data.ram.map(item => `ID ${item.id}`),
        datasets: [
            {
                label: 'RAM Used (%)',
                data: data.ram.map(item => item.used),
                backgroundColor: 'rgba(255, 206, 86, 0.5)',
            },
            {
                label: 'RAM Free (%)',
                data: data.ram.map(item => item.free),
                backgroundColor: 'rgba(75, 192, 192, 0.5)',
            }
        ],
    };

    return (


        <><h1>CPU Usage</h1><Bar options={{ responsive: true }} data={cpuData} /><h1>RAM Usage</h1><Bar options={{ responsive: true }} data={ramData} /></>


    );
};

export default ChartComponent;
