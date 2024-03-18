import { React, useEffect, useState } from 'react';
import { graphviz } from 'd3-graphviz';

const ProcessSimulator = () => {
    const [pid, setPid] = useState(0);
    const [state, setState] = useState('New');
    const [dot, setDot] = useState('');


    useEffect(() => {
        let currentdot = "digraph {"
        switch (state) {

            case 'Ready':
                currentdot += `d${pid} [shape=record, style=filled, color=blue];`
                break;
            case 'Running':
                currentdot += `d${pid} [shape=record, style=filled, color=green];`
                break;
            case 'Terminated':
                currentdot += `d${pid} [shape=record, style=filled, color=red];`
                break;
            default:
                currentdot += `d${pid} [shape=record, style=filled, color=lightgrey];`
                break;
        }

        switch (state) {

            case 'Ready':
                currentdot += `Running -> Ready  Running [color="blue"];`
                break;
            case "Running":
                currentdot += `New -> Ready -> Running [color="green"];`
                break;
            case "Terminated":
                currentdot += `Running -> Terminated [color="red"]`
                break;
            default:
                currentdot += `New -> ${pid};`
                break;
        }
        currentdot += "}"

        setDot(currentdot);
    }, [state, pid]);

    const createNewProcess = () => {
        // Aquí generas un nuevo PID y estableces el estado inicial del proceso
        // const newPid = Math.floor(Math.random() * 100000); // Un ejemplo simple de generación de PID
        // setPid(newPid);
        fetch("/start", {
            method: "GET"
        }).then((response) => {
            setPid(parseInt(response));
            setState('Running');
        });
        // setState('Running');
    };

    const killProcess = () => {
        // setState('Terminated');
        fetch("/kill?pid="+pid, {
            method: "GET"
        }).then((response) => {
            setPid(0);
            setState('Terminated');
        });
    };

    const stopProcess = () => {
        // setState('Ready');
        fetch("/stop?pid="+pid, {
            method: "GET"
        }).then((response) => {
            setState('Ready');
        });
    };

    const resumeProcess = () => {
        // setState('Running');
        fetch("/resume?pid="+pid, {
            method: "GET"
        }).then((response) => {
            setState('Running');
        });
    };


    useEffect(() => {
        console.log(dot);
        graphviz("#graph2").renderDot(dot);
    }, [dot]);

    return (

        <div>
            <h1>Proceso {pid}</h1>
            <div id="graph2"></div>
            <button onClick={createNewProcess}>New</button>
            <button onClick={killProcess} disabled={!pid}>Kill</button>
            <button onClick={stopProcess} disabled={!pid || state !== 'Running'}>Stop</button>
            <button onClick={resumeProcess} disabled={!pid || state !== 'Ready'}>Resume</button>
        </div>
    );
}

export default ProcessSimulator;
