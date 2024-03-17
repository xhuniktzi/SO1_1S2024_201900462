import { graphviz } from 'd3-graphviz'
import { React, useEffect, useState } from 'react'
import { Chart, PieController, ArcElement, Tooltip, Legend } from 'chart.js';
Chart.register(PieController, ArcElement, Tooltip, Legend);

const recursiveGenerateDot = (data) => {
    if (!data) {
        return ''
    }

    let dot = `node${data.PID} [label="{PID: ${data.PID} | Name: ${data.Name} }"];\n`

    if (data.Children) {
        data.Children.forEach(child => {
            dot += `node${data.PID} -> node${child.PID};\n`
            dot += recursiveGenerateDot(child)
        })
    }

    return dot
}

const recursiveFindPidLayer0 = (pid, data) => {

    if (!data) {
        return
    }

    for (let i = 0; i < data.length; i++) {
        if (data[i].PID == pid) {

            return data[i]
        } else {
            recursiveFindPidLayer0(pid, data[i].Children)
        }
    }
}

const TreeProcess = () => {
    const [data, setData] = useState([])
    const [dot, setDot] = useState('')



    const fetchData = async () => {
        const response = await fetch('/cpu')
        setData(await response.json())

        const pid = document.getElementById('pid').value

        const result = recursiveFindPidLayer0(pid, data.Processes)
        setDot(recursiveGenerateDot(result))
    }

    useEffect(() => {
        graphviz("#graph")
            .renderDot(`digraph { node [shape=record];\n ${dot} \n}`)
    }, [dot])

    return (
        <>
            <label htmlFor="pid">PID: </label>
            <input type="number" id="pid" name="pid" />
            <button type="button" onClick={fetchData}>Buscar</button>
            <div id="graph" />
        </>
    )
}

export default TreeProcess;